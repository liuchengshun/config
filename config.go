package config

import "sync"

var CONF *conf

var once sync.Once

func init() {
	once.Do(func() {
		CONF = newConf()
	})
}
