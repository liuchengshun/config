package config

import (
	"strconv"
	"sync"
)

// the default value when get value by key failed.
const (
	defaultResultString = ""
	defaultResultBool   = false
	defaultResultInt    = 0
)

type confGroup struct {
	name     string
	elements map[string]string
	mu       sync.RWMutex
}

func newConfGroup(name string) *confGroup {
	return &confGroup{
		name:     name,
		elements: make(map[string]string),
	}
}

func (g *confGroup) set(key, value string) {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.elements[key] = value
}

func (g *confGroup) getString(key string) string {
	return g.get(key)
}

func (g *confGroup) getInt(key string) int {
	v := g.get(key)
	if v == "" {
		return defaultResultInt
	}

	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return defaultResultInt
	}
	return int(i)
}

func (g *confGroup) getBool(key string) bool {
	v := g.get(key)
	if v == "" {
		return defaultResultBool
	}

	b, err := strconv.ParseBool(v)
	if err != nil {
		return defaultResultBool
	}
	return b
}

func (g *confGroup) get(key string) string {
	g.mu.RLock()
	defer g.mu.RUnlock()

	return g.elements[key]
}
