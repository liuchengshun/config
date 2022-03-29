package config

import "sync"

type configdata struct {
	elements map[string]string
	mu       sync.RWMutex
}

func (c *configdata) set(key, value string) {
	
}
