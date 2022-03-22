package config

func init() {
	CONF = &config{}
}

var CONF *config

type config struct {
	SetReader

	fileType fileType
}

func (c *config) LoadConfig(path string) {
	ftype := parseFileType(path)

}

func (c *config) loadKeyValueConfig(path string) {

}
