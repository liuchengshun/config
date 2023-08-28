package config

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
)

type Conf struct {
	confGroups    []*confGroup
	defaultGroups []*Section
	errors        *confErrors
}

func newConf() *Conf {
	return &Conf{
		confGroups:    make([]*confGroup, 0),
		defaultGroups: make([]*Section, 0),
		errors:        newConfError(),
	}
}

func (c *Conf) LoadConfiguration(filePath string) error {
	ext := filepath.Ext(filePath)
	if ext != ".conf" {
		return fmt.Errorf("does not support file %s with %s", filePath, ext)
	}

	return c.loadConfiguration(filePath)
}

func (c *Conf) loadConfiguration(filePath string) error {
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

		confline := newConfLine(text)
		lineT, err := confline.parse()
		if err != nil {
			c.errors.appendError(err)
			continue
		}

		switch lineT {
		case lineEmpty:
			continue
		case lineConfGroup:
			name := confline.confGroupName()
			if name != "" {
				group = c.getConfGroup(name)
			}
		case lineKeyValue:
			key, value := confline.keyValue()
			if key != "" && group != nil {
				group.set(key, value)
			}
		}
	}

	if scanner.Err() != nil {
		return err
	}

	if !c.errors.isNil() {
		return c.errors
	}

	return nil
}

func (c *Conf) getConfGroup(name string) *confGroup {
	for _, g := range c.confGroups {
		if g.name == name {
			return g
		}
	}
	g := newConfGroup(name)
	c.confGroups = append(c.confGroups, g)
	return g
}

func (c *Conf) RegisterGroup(g *Section) {
	for _, gro := range c.defaultGroups {
		if gro.name == g.name {
			gro.copy(g)
			return
		}
	}

	ng := NewSection(g.name)
	ng.copy(g)
	c.defaultGroups = append(c.defaultGroups, ng)
}

// if get failed, GetString will return the empty string.
func (c *Conf) GetString(group, key string) string {
	// get from groups.
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

// if get failed, GetBool will return false.
func (c *Conf) GetBool(group, key string) bool {
	// get from groups.
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

// if get failed, GetInt will return -1.
func (c *Conf) GetInt(group, key string) int {
	// get from groups.
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
