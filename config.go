package config

import (
	"fmt"
	"path/filepath"
	"sync"
)

var CONF *conf

var once sync.Once

func LoadConfiguration(filePath string) error {
	ext := parseFileExtension(filePath)
	if ext != configFileConfExt {
		return fmt.Errorf("does not support the extension name %s, only support .conf", ext)
	}

	once.Do(func() {
		CONF = newConf(filePath)
	})
	return CONF.loadConfiguration()
}

const configFileConfExt = ".conf"

func parseFileExtension(filePath string) string {
	ext := filepath.Ext(filePath)
	if ext == configFileConfExt {
		return configFileConfExt
	}
	return ""
}
