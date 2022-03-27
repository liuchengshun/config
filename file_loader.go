package config

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type fileLoader interface {
	confLoader
}

type confLoader interface {
	loadCONF() map[string]map[string]string
}

type fileLoad struct {
	path string
}

func newFileLoader(path string) fileLoader {
	return fileLoad{
		path: path,
	}
}

const (
	fileTypeCONF = "file_type_conf"
	fileTypeJSON = "file_type_json"
)

func (l fileLoad) filetype() string {
	switch filepath.Ext(l.path) {
	case ".conf":
		return fileTypeCONF
	case ".json":
		return fileTypeJSON
	}
	return ""
}

func (l fileLoad) loadCONF() map[string]map[string]string {
	configs := make(map[string]map[string]string)

	filebytes, err := os.ReadFile(l.path)
	if err != nil {
		panic(err)
	}
	rendered := bytes.NewReader(filebytes)

	var group string
	scanner := bufio.NewScanner(rendered)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		text := scanner.Text()

		if strings.HasPrefix(text, "[") && strings.HasSuffix(text, "]") && len(text) >= 3 {
			group = text[1 : len(text)-2]
			_, ok := configs[group]
			if !ok {
				configs[group] = make(map[string]string)
			}
			continue
		}

		parts := strings.Split(text, "=")
		if len(parts) != 2 {
			panic(fmt.Sprintf("the config message is error, does not support the format: %s", text))
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		if group != "" {
			v := strings.Trim(value, "\"")
			configs[group][key] = v
		}
	}

	if scanner.Err() != nil {
		panic(err)
	}

	return configs
}
