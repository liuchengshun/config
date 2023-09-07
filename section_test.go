package inifile

import (
	"reflect"
	"testing"
)

func TestSectionSet(t *testing.T) {
	tests := []struct {
		key   string
		value interface{}
	}{
		{"hello", "world"},
	}
	for _, tt := range tests {
		sec := NewSection("test")
		actual := sec.Set(tt.key, tt.value).Get(tt.key)
		if !reflect.DeepEqual(actual, tt.value) {
			t.Fatalf("Set item failed, the expect value is %v, but got %v", tt.value, actual)
		}
	}
}
