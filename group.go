package config

type group struct {
	name    string
	configs map[string]interface{}
}

func NewGroup(name string) *group {
	return &group{
		name:    name,
		configs: make(map[string]interface{}),
	}
}

func (g *group) SetString(key, v string) {
	g.configs[key] = v
}

func (g *group) SetBool(key string, v bool) {
	g.configs[key] = v
}

func (g *group) SetInt(key string, v int) {
	g.configs[key] = v
}

func (g *group) copy(src *group) {
	if g.name != src.name {
		return
	}
	for k, v := range src.configs {
		if _, ok := g.configs[k]; !ok {
			g.configs[k] = v
		}
	}
}
