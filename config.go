package config

import (
	"fmt"
	"path/filepath"
	"sync"
)

var CONF *conf

var once sync.Once

// load configuration by file path, and only the first call is valid.
func LoadConfiguration(filePath string) error {
	ext := parseFileExtension(filePath)
	if ext != configFileConfExt {
		return fmt.Errorf("does not support the extension name %s, only support .conf", ext)
	}

	var err error
	once.Do(func() {
		CONF = newConf(filePath)
		err = CONF.loadConfiguration()
	})
	return err
}

const configFileConfExt = ".conf"

func parseFileExtension(filePath string) string {
	ext := filepath.Ext(filePath)
	if ext == configFileConfExt {
		return configFileConfExt
	}
	return ""
}
