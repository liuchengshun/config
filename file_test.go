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
