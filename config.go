package config

import "sync"

var conf *Conf

var once sync.Once

func CONF() *Conf {
	once.Do(func() {
		conf = newConf()
	})
	return conf
}
