package config

type conf struct {
	confLoader

	groups []*Group
}

func newConf(path string) *conf {
	return &conf{
		groups:     make([]*Group, 0),
		confLoader: newFileLoader(path),
	}
}

func (c *conf) initGroups() {
	configs := c.loadCONF()

	for gname, values := range configs {
		group := c.getGroup(gname)
		for key, value := range values {
			group.set(key, value)
		}
	}
}

func (c *conf) getGroup(name string) *Group {
	for _, g := range c.groups {
		if name == g.name {
			return g
		}
	}
	g := NewGroup(name)
	c.groups = append(c.groups, g)
	return g
}

func (c *conf) RegisterGroup(g *Group) {
	c.registerGroup(g)
}

func (c *conf) registerGroup(g *Group) {
	for _, gro := range c.groups {
		if gro.name == g.name {
			gro.copy(g)
			return
		}
	}
	c.groups = append(c.groups, g)
}

// if read failed, will return the empty string.
func (c *conf) ReadString(group, key string) string {
	for _, g := range c.groups {
		if g.name == group {
			return g.readString(key)
		}
	}
	return defaultString
}

// if read failed, ReadBool will return false.
func (c *conf) ReadBool(group, key string) bool {
	for _, g := range c.groups {
		if g.name == group {
			return g.readBool(key)
		}
	}
	return defaultBool
}

// if read failed, ReadInt will return -1.
func (c *conf) ReadInt(group, key string) int64 {
	for _, g := range c.groups {
		if g.name == group {
			return g.readInt(key)
		}
	}
	return defaultInt
}
