package config

import "sync"

var CONF *conf

var once sync.Once

func LoadConfiguration(filePath string) error {
	once.Do(func() {
		CONF = newConf(filePath)
	})
	return CONF.loadConfiguration()
}
