package config

import (
	"fmt"
	"strings"
)

const (
	lineEmpty = iota
	lineConfGroup
	lineKeyValue
	lineErr
)

type confLine struct {
	text string

	lineT int

	key       string
	value     string
	groupName string
}

func newConfLine(text string) *confLine {
	return &confLine{text: text}
}

func (l *confLine) parse() (lineType int, err error) {
	text := l.trim()

	// is empty string.
	if text == "" {
		return lineEmpty, nil
	}

	// parse key&value
	parts := []string{}
	switch {
	case strings.Contains(text, " = "):
		parts = strings.Split(text, " = ")
	case strings.Contains(text, " ="):
		parts = strings.Split(text, " =")
	}
	if len(parts) == 2 {
		l.lineT = lineKeyValue
		l.key = parts[0]
		l.value = strings.Trim(parts[1], "\"")
		return l.lineT, nil
	} else if len(parts) != 2 && len(parts) > 0 {
		l.lineT = lineErr
		return lineErr, fmt.Errorf("does not support the config line: %s", l.text)
	}

	// parse conf group name.
	if len(text) >= 3 && strings.HasPrefix(text, "[") && strings.HasSuffix(text, "]") {
		l.groupName = text[1 : len(text)-1]
		l.lineT = lineConfGroup
		return lineConfGroup, nil
	}

	l.lineT = lineErr
	return lineErr, fmt.Errorf("does not support the config line: %s", l.text)
}

func (l *confLine) trim() string {
	text := strings.TrimSpace(l.text)
	for i, v := range text {
		if v == '#' {
			// if index is 0.
			if i == 0 {
				return ""
			}

			// if index is not 0.
			if text[i-1] == ' ' {
				return strings.TrimSpace(text[:i-1])
			}
		}
	}
	return text
}

func (l *confLine) keyValue() (key, value string) {
	if l.lineT == lineKeyValue {
		key = l.key
		value = l.value
		return
	}
	return
}

func (l *confLine) confGroupName() string {
	if l.lineT == lineConfGroup {
		return l.groupName
	}
	return ""
}
