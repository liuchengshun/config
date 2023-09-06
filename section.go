package inifile

import (
	"reflect"
	"strconv"
	"sync"
)

type Section struct {
	name  string
	items map[string]string
	mu    sync.RWMutex
}

func NewSection(name string) *Section {
	return &Section{
		name:  name,
		items: make(map[string]string),
	}
}

// Set sets the item when value is integer, float, map[string]string, map[string]interface{}
// And other types like func, channel, slice ... would not be set
func (s *Section) Set(key string, value interface{}) {
	if value == nil {
		return
	}
	if val := s.parseValue(value); val != "" {
		s.setString(key, val)
		return
	}
	// set the item when value type is map[string]string
	if mapval, ok := value.(map[string]string); ok {
		for key, value := range mapval {
			s.setString(key, value)
		}
	}
	// set the item when value type is map[string]interface{}
	if mapval, ok := value.(map[string]interface{}); ok {
		for key, value := range mapval {
			if val := s.parseValue(value); val != "" {
				s.setString(key, val)
			}
		}
	}
}

func (s *Section) setString(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items[key] = value
}

func (s *Section) parseValue(value interface{}) string {
	if value == nil {
		return ""
	}
	vv := reflect.ValueOf(value)
	k := vv.Kind()
	var val string
	switch {
	case k == reflect.String:
		val = vv.String()
	case k == reflect.Bool:
		if vv.Bool() {
			val = "true"
		} else {
			val = "false"
		}
	case k >= reflect.Int && k <= reflect.Uint64:
		val = strconv.FormatInt(vv.Int(), 10)
	case k == reflect.Float32:
		val = strconv.FormatFloat(vv.Float(), 'E', -1, 32)
	case k == reflect.Float64:
		val = strconv.FormatFloat(vv.Float(), 'E', -1, 64)
	}
	return val
}

func (s *Section) Get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.items[key]
}

func (s *Section) Range(op func(key, value string) bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for key, value := range s.items {
		if ok := op(key, value); !ok {
			return
		}
	}
	return
}
