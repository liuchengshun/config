package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type IniFile struct {
	sections []Section
}

func NewIniFile() *IniFile {
	return &IniFile{
		sections: make([]Section, 0),
	}
}

func LoadIniFile(filePath string) (*IniFile, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("read file %s failed: %v", filePath, err)
	}
	iniFile := NewIniFile()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// skip blank and single line comment
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

	}
	if scanner.Err() != nil {
		return nil, fmt.Errorf("parse file content failed: %v", err)
	}
	return iniFile, nil
}

func parseSectionName(line string) (string, bool) {
	if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
		name := strings.TrimPrefix(line, "[")
		name = strings.TrimSuffix(name, "]")
		return name, true
	}
	return "", false
}

func parseKeyValue(line string) (key, value string, ok bool) {
	parts := strings.SplitN(line, "=", 2)
	if len(parts) != 2 {
		return "", "", false
	}
	return parts[0], parts[1], true
}

func cleanLine(line string) string {
	hasAnnotation, hasSingleQuota, hasDoubleQuota := false, false, false
	interceptIndex := 0
	for i := 0; i < len(line); i++ {
		v := line[i]
		if v == '\'' {
			if hasSingleQuota {
				hasDoubleQuota = false
				hasAnnotation = false
				hasSingleQuota = false
				continue
			} else {
				hasSingleQuota = true
			}
		}
		if v == '"' {
			if hasDoubleQuota {
				hasDoubleQuota = false
				hasSingleQuota = false
				hasAnnotation = false
				continue
			} else {
				hasDoubleQuota = true
			}
		}
		if v == '#' {
			if !hasSingleQuota && !hasDoubleQuota {
				return line[i+1:]
			}
			if hasSingleQuota || hasDoubleQuota {
				interceptIndex = i // 这里需要优化！！
			}
		}
	}
}
