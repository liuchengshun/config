package config

import (
	"bytes"
	"os"
)

type keyValConfig struct {
	configs map[string]map[string]interface{}
}

func (c *keyValConfig) loadConfig(path string) {
	filebytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	buf := bytes.NewReader(filebytes)
	for {
		
	}
}
