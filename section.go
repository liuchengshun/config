package inifile

import (
	"fmt"
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
// there will be panic, if other types like func, channel, slice as value
func (s *Section) Set(key string, value interface{}) *Section {
	if value == nil {
		return s
	}
	if val := parseValue(value); val != "" {
		s.setString(key, val)
		return s
	}
	return s
}

func (s *Section) setString(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items[key] = value
}

// SetMap cloud set the item for section by map[string]string and map[string]interface{}
func (s *Section) SetMap(value interface{}) *Section {
	switch v := value.(type) {
	case map[string]string:
		for key, value := range v {
			s.setString(key, value)
		}
	case map[string]interface{}:
		for key, value := range v {
			vv := parseValue(value)
			s.setString(key, vv)
		}
	}
	return s
}

func parseValue(value interface{}) string {
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
	default:
		panic(fmt.Sprintf("section items does not support the value of %v", vv.Type()))
	}
	return val
}

// Get is like GetVal, but not return the bool value.
func (s *Section) Get(key string) string {
	v, _ := s.GetVal(key)
	return v
}

// GetVal gets the value by key, if the value is not exist, return empty string and false.
func (s *Section) GetVal(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.items[key]
	return v, ok
}

// Range cloud visit the every item of section,
// the argument op clound support the operation for item, Note that returned false will skip all other items.
func (s *Section) Range(op func(key, value string) bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for key, value := range s.items {
		if ok := op(key, value); !ok {
			return
		}
	}
}
