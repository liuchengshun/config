package config

var CONF configer

type configer interface {
	LoadConfiguration(path string)

	ReadString(group, key string) string
	ReadInt(group, key string) int64
	ReadBool(group, key string) bool

	RegisterGroup(g *Group)

	initGroups()
}

func LoadConfiguration(path string) {

}
