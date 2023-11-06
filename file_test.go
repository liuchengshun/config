package inifile

import (
	"testing"
)

func TestCleanLine(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{"[hello] # aaao", "[hello]"},
		{"# helesdf", ""},
		{"aaa# helesdf", "aaa"},
		{"a#aa# helesdf", "a"},
		{`"#dfasdf"#'dff'`, `"#dfasdf"`},
	}
	for _, tt := range tests {
		actual := cleanComment(tt.input)
		if actual != tt.expect {
			t.Fatalf("expect %s length = %v, but got %s, length = %v", tt.expect, len(tt.expect), actual, len(actual))
		}
	}
}

func TestParseKeyValue(t *testing.T) {
	tests := []struct {
		input         string
		expectedKey   string
		expectedValue string
		expectedOk    bool
	}{
		{"hellin=your my luck", "hellin", "your my luck", true},
		{"hellin = your my luck", "hellin", "your my luck", true},
		{"  hellin   =   your my luck", "hellin", "your my luck", true},
		{" hellin  = your my luck", "hellin", "your my luck", true},
		{" hel lin  = your my luck", "hel lin", "your my luck", true},
	}
	for _, tt := range tests {
		key, value, ok := parseKeyValue(tt.input)
		if ok != tt.expectedOk {
			t.Fatalf("expect parse %v, but got %v", tt.expectedOk, ok)
		}
		if !tt.expectedOk {
			continue
		}
		if tt.expectedKey != key {
			t.Fatalf("expect key %v, but got %v", tt.expectedKey, tt.expectedKey)
		}
		if tt.expectedValue != value {
			t.Fatalf("expect value %v, but got %v", tt.expectedValue, value)
		}
	}
}

func TestParseSectionName(t *testing.T) {
	tests := []struct {
		line        string
		expect      string
		parseResult bool
	}{
		{"[hello]", "hello", true},
		{" [hello] ", "hello", true},
		{"   [he llo]", "he llo", true},
	}
	for _, tt := range tests {
		line := cleanComment(tt.line)
		actual, ok := parseSectionName(line)
		if ok != tt.parseResult {
			t.Fatalf("expect parse result = %v, but got %v", tt.parseResult, ok)
		}
		if actual != tt.expect {
			t.Fatalf("expect result = %v, but got %v", tt.expect, actual)
		}
	}
}

func TestLoadIniFile(t *testing.T) {
	file, err := LoadIniFile("./test/testdata/config.conf")
	if err != nil {
		t.Fatalf("load ini file failed: %v", err)
	}
	serverSection, ok := file.GetSection("server")
	if !ok || serverSection == nil {
		t.Fatalf("[loadIniFile] get server section failed")
	}
}
