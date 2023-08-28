package config

import (
	"strconv"
	"sync"
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

func (g *confGroup) getString(key string) (string, bool) {
	return g.get(key)
}

func (g *confGroup) getInt(key string) (int, bool) {
	val, ok := g.get(key)
	if !ok {
		return 0, false
	}

	v, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0, false
	}
	return int(v), true
}

// the value would be 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False.
func (g *confGroup) getBool(key string) (v bool, ok bool) {
	val, ok := g.get(key)
	if !ok {
		return false, false
	}

	b, err := strconv.ParseBool(val)
	if err != nil {
		return false, false
	}
	return b, true
}

func (g *confGroup) get(key string) (string, bool) {
	g.mu.RLock()
	defer g.mu.RUnlock()

	v, ok := g.elements[key]
	return v, ok
}
