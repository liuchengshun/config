package inifile

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type IniFile struct {
	sections []*Section
}

func NewIniFile() *IniFile {
	return &IniFile{
		sections: make([]*Section, 0),
	}
}

func LoadIniFile(filePath string) (*IniFile, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("read file %s failed: %v", filePath, err)
	}
	iniFile := NewIniFile()
	scanner := bufio.NewScanner(file)
	var curSection *Section
	var isDupLatestSec bool
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// skip blank and single line comment
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		line = cleanComment(line)

		secName, ok := parseSectionName(line)
		if ok {
			if curSection != nil {
				// the same section appears consecutively
				if curSection.name == secName {
					continue
				}
				// the current section is registered duplicately in the same ini file.
				if sec, ok := iniFile.GetSection(secName); ok {
					curSection = sec
					isDupLatestSec = true
					continue
				}
				isDupLatestSec = false
				iniFile.MergeSection(curSection)
				curSection = NewSection(secName)
			} else {
				curSection = NewSection(secName)
				continue
			}
		}

		key, value, ok := parseKeyValue(line)
		if ok {
			if curSection == nil {
				return nil, fmt.Errorf("missing the section before set %s=%s", key, value)
			}
			curSection.Set(key, value)
			continue
		}

		return nil, fmt.Errorf("the line is unknow = %v", line)
	}
	if scanner.Err() != nil {
		return nil, fmt.Errorf("parse file content failed: %v", err)
	}
	if !isDupLatestSec {
		iniFile.MergeSection(curSection)
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

func cleanComment(line string) string {
	hasSingleQuota, hasDoubleQuota := false, false
	interceptIndex := -1
	for i := 0; i < len(line); i++ {
		v := line[i]
		if v == '\'' {
			if hasSingleQuota {
				hasDoubleQuota = false
				hasSingleQuota = false
				interceptIndex = -1
				continue
			} else {
				hasSingleQuota = true
			}
		}
		if v == '"' {
			if hasDoubleQuota {
				hasDoubleQuota = false
				hasSingleQuota = false
				interceptIndex = -1
				continue
			} else {
				hasDoubleQuota = true
			}
		}
		if v == '#' {
			if !hasSingleQuota && !hasDoubleQuota {
				line = line[:i]
				break
			}
			interceptIndex = i
		}
	}
	if interceptIndex != -1 {
		line = line[:interceptIndex]
	}
	return strings.TrimSpace(line)
}

func (f *IniFile) Section(secname string) *Section {
	sec, _ := f.GetSection(secname)
	return sec
}

func (f *IniFile) GetSection(secname string) (*Section, bool) {
	for _, sec := range f.sections {
		if sec.name == secname {
			return sec, true
		}
	}
	return NewSection(secname), false
}

func (f *IniFile) AddSection(secname string) *Section {
	for _, s := range f.sections {
		if s.name == secname {
			return s
		}
	}
	sec := NewSection(secname)
	f.sections = append(f.sections, sec)
	return sec
}

func (f *IniFile) MergeSection(sec *Section) {
	if sec == nil {
		return
	}
	var dupSec *Section
	for i := 0; i < len(f.sections); i++ {
		if f.sections[i].name == sec.name {
			dupSec = f.sections[i]
		}
	}
	if dupSec != nil {
		sec.Range(func(key, value string) bool {
			dupSec.Set(key, value)
			return true
		})
	}
}

func (f *IniFile) LoadToFile() {
	
}
