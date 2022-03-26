package config

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

type conf struct {
	groups []*group
}

func newConf() *conf {
	return &conf{
		groups: make([]*group, 0),
	}
}

func (c *conf) loadConfig(path string) {
	filebytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	rendered := bytes.NewReader(filebytes)

	var group *group
	scanner := bufio.NewScanner(rendered)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		text := scanner.Text()

		if strings.HasPrefix(text, "[") && strings.HasSuffix(text, "]") && len(text) >= 3 {
			name := text[1 : len(text)-2]
			group = c.getGroup(name)
			continue
		}

		parts := strings.Split(text, "=")
		if len(parts) != 2 {
			panic(fmt.Sprintf("the config message is error, does not support the format: %s", text))
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		if group != nil {
			group.SetString(key, value)
		}
	}

	if scanner.Err() != nil {
		panic(err)
	}
}

func (c *conf) getGroup(name string) *group {
	for _, g := range c.groups {
		if name == g.name {
			return g
		}
	}
	g := NewGroup(name)
	c.groups = append(c.groups, g)
	return g
}

func (c *conf) registerGroup(g *group) *group {
	
}
