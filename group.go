package config

import "sync"

type Group struct {
	name     string
	elements map[string]interface{}
	mu       sync.RWMutex
}

func NewGroup(name string) *Group {
	return &Group{
		name:     name,
		elements: make(map[string]interface{}),
	}
}

func (g *Group) SetString(key, v string) {
	g.set(key, v)
}

func (g *Group) SetBool(key string, v bool) {
	g.set(key, v)
}

func (g *Group) SetInt(key string, v int) {
	g.set(key, v)
}

func (g *Group) set(key string, v interface{}) {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.elements[key] = v
}

func (g *Group) getString(key string) (v string, ok bool) {
	val := g.get(key)
	if val != nil {
		v, ok = val.(string)
		return
	}
	return
}

func (g *Group) getBool(key string) (v bool, ok bool) {
	val := g.get(key)
	if val != nil {
		v, ok = val.(bool)
		return
	}
	return
}

func (g *Group) getInt(key string) int {
	v := g.get(key)
	if v != nil {
		i, ok := v.(int)
		if ok {
			return i
		}
	}
	return defaultResultInt
}

func (g *Group) get(key string) interface{} {
	g.mu.RLock()
	defer g.mu.RUnlock()

	return g.elements[key]
}

func (g *Group) copy(gro *Group) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.name != gro.name {
		return
	}

	for k, v := range gro.elements {
		_, ok := g.elements[k]
		if !ok {
			g.elements[k] = v
		}
	}
}
