package config

import (
	"strconv"
	"sync"
)

type Group struct {
	name    string
	configs map[string]interface{}
	mu      sync.RWMutex
}

func NewGroup(name string) *Group {
	return &Group{
		name:    name,
		configs: make(map[string]interface{}),
	}
}

func (g *Group) SetString(key, v string) {
	g.set(key, v)
}

func (g *Group) SetBool(key string, v bool) {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.configs[key] = v
}

func (g *Group) SetInt(key string, v int) {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.configs[key] = v
}

func (g *Group) set(key, v string) {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.configs[key] = v
}

const (
	defaultString = ""
	defaultBool   = false
	defaultInt    = -1
)

func (g *Group) getString(key string) string {
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

func (g *Group) getBool(key string) bool {
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

func (g *Group) getInt(key string) int64 {
	g.mu.RLock()
	defer g.mu.RUnlock()

	val, ok := g.configs[key]
	if ok {
		switch v := val.(type) {
		case int64:
			return v
		case string:
			i, err := strconv.ParseInt(v, 10, 64)
			if err == nil {
				return i
			}
		}
	}
	return defaultInt
}

func (g *Group) copy(src *Group) {
	g.mu.RLock()
	defer g.mu.RUnlock()

	if g.name != src.name {
		return
	}
	for k, v := range src.configs {
		if _, ok := g.configs[k]; !ok {
			g.configs[k] = v
		}
	}
}
