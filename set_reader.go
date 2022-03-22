package config

type SetReader interface {
	Set(group, key string)

	GetString(group, key string) string
	GetInt(group, key string) int
	GetBool(group, key string) bool
}
