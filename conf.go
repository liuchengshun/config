package config

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type conf struct {
	confGroups    []*confGroup
	defaultGroups []*Group
}

func newConf() *conf {
	return &conf{
		confGroups:    make([]*confGroup, 0),
		defaultGroups: make([]*Group, 0),
	}
}

func (c *conf) LoadConfiguration(filePath string) error {
	ext := filepath.Ext(filePath)
	if ext != ".conf" {
		return fmt.Errorf("does not support file %s with %s", filePath, ext)
	}

	return c.loadConfiguration(filePath)
}

func (c *conf) loadConfiguration(filePath string) error {
	filebytes, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("read file data failed: %v", err)
	}
	rendered := bytes.NewReader(filebytes)

	var group *confGroup
	scanner := bufio.NewScanner(rendered)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		text := scanner.Text()

		// skip empty line.
		if text == "" {
			continue
		}

		// skip annotation
		if strings.HasPrefix(text, "#") {
			continue
		}

		// parse [group] line.
		if strings.HasPrefix(text, "[") && strings.HasSuffix(text, "]") && len(text) >= 3 {
			name := text[1 : len(text)-1]
			group = c.getConfGroup(name)
			continue
		}

		// parse key=value line.
		parts := strings.SplitN(text, " = ", 2)
		if len(parts) != 2 {
			return fmt.Errorf("the config message is error, does not support the format: %s", text)
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		if group != nil {
			v := strings.Trim(value, "\"")
			group.set(key, v)
		}
	}

	if scanner.Err() != nil {
		return err
	}

	return nil
}

func (c *conf) getConfGroup(name string) *confGroup {
	for _, g := range c.confGroups {
		if g.name == name {
			return g
		}
	}
	g := newConfGroup(name)
	c.confGroups = append(c.confGroups, g)
	return g
}

func (c *conf) RegisterGroup(g *Group) {
	for _, gro := range c.defaultGroups {
		if gro.name == g.name {
			gro.copy(g)
			return
		}
	}

	ng := NewGroup(g.name)
	ng.copy(g)
	c.defaultGroups = append(c.defaultGroups, ng)
}

// if get failed, will return the empty string.
func (c *conf) GetString(group, key string) string {
	// get from conf group.
	for _, g := range c.confGroups {
		if g.name == group {
			v, ok := g.getString(key)
			if !ok {
				break
			}
			return v
		}
	}

	// get from default group.
	for _, g := range c.defaultGroups {
		if g.name == group {
			v, ok := g.getString(key)
			if !ok {
				break
			}
			return v
		}
	}

	return ""
}

// if get failed, getBool will return false.
func (c *conf) GetBool(group, key string) bool {
	// get from conf groups.
	for _, g := range c.confGroups {
		if g.name == group {
			v, ok := g.getBool(key)
			if !ok {
				break
			}
			return v
		}
	}

	// get from default groups.
	for _, g := range c.defaultGroups {
		if g.name == group {
			v, ok := g.getBool(key)
			if !ok {
				break
			}
			return v
		}
	}
	return false
}

// if get failed, getInt will return -1.
func (c *conf) GetInt(group, key string) int {
	// get from conf groups.
	for _, g := range c.confGroups {
		if g.name == group {
			v, ok := g.getInt(key)
			if !ok {
				break
			}
			return v
		}
	}

	// get from default groups.
	for _, g := range c.defaultGroups {
		if g.name == group {
			v, ok := g.getInt(key)
			if !ok {
				break
			}
			return v
		}
	}

	return 0
}
