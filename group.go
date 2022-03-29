package config

import (
	"strconv"
	"sync"
)

type ConfGroup struct {
	name    string
	configs map[string]interface{}
	mu      sync.RWMutex
}

func NewConfGroup(name string) *ConfGroup {
	return &ConfGroup{
		name:    name,
		configs: make(map[string]interface{}),
	}
}

func (g *ConfGroup) SetString(key, v string) {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.set(key, v)
}

func (g *ConfGroup) SetBool(key string, v bool) {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.set(key, v)
}

func (g *ConfGroup) SetInt(key string, v int) {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.set(key, v)
}

func (g *ConfGroup) set(key string, v interface{}) {
	g.configs[key] = v
}

const (
	defaultString = ""
	defaultBool   = false
	defaultInt    = -1
)

func (g *ConfGroup) getString(key string) string {
	g.mu.RLock()
	defer g.mu.RUnlock()

	v, ok := g.configs[key]
	if ok {
		s, ok := v.(string)
		if ok {
			return s
		}
	}
	return defaultString
}

func (g *ConfGroup) getBool(key string) bool {
	g.mu.RLock()
	defer g.mu.RUnlock()

	val, ok := g.configs[key]
	if ok {
		b, ok := val.(bool)
		if ok {
			return b
		}
		switch v := val.(type) {
		case bool:
			return v
		case string:
			b, err := strconv.ParseBool(v)
			if err != nil {
				return defaultBool
			}
			return b
		}
	}
	return defaultBool
}

func (g *ConfGroup) getInt(key string) int {
	g.mu.RLock()
	defer g.mu.RUnlock()

	val, ok := g.configs[key]
	if ok {
		switch v := val.(type) {
		case int:
			return v
		case string:
			i, err := strconv.ParseInt(v, 10, 64)
			if err == nil {
				return int(i)
			}
		}
	}
	return defaultInt
}

func (g *ConfGroup) copy(src *ConfGroup) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.name != src.name {
		return
	}
	for k, v := range src.configs {
		if _, ok := g.configs[k]; !ok {
			g.configs[k] = v
		}
	}
}

func (g *ConfGroup) clone() *ConfGroup {
	ng := NewConfGroup(g.name)
	for k, v := range g.configs {
		ng.set(k, v)
	}
	return ng
}
