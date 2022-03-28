package test

import (
	"config"
	"reflect"
	"testing"
)

func TestLoadConfiguration(t *testing.T) {
	
}

func TestCONF(t *testing.T) {
	err := config.LoadConfiguration("./testdata/config.conf")
	if err != nil {
		t.Fatal(err)
	}

	CONF := config.CONF

	tests := []struct {
		name      string
		valueType string
		group     string
		key       string
		expected  interface{}
	}{
		{"server group", "string", "server", "host", "127.0.0.1"},
		{"server group", "string", "server", "port", "8080"},
		{"company group", "string", "company", "name", "bukaopu"},
		{"company group", "int", "company", "age", int64(80)},
		{"company group", "string", "company", "business", "write_code"},
		{"company group", "string", "company", "usedlang", "golang"},
		{"language group", "string", "language", "name", "golang"},
		{"language group", "string", "language", "maxcot", "mouse"},
		{"language group", "string", "language", "version", "1.17.7"},
		{"language group", "bool", "language", "production", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var actual interface{}
			switch tt.valueType {
			case "string":
				actual = CONF.GetString(tt.group, tt.key)
			case "int":
				actual = CONF.GetInt(tt.group, tt.key)
			case "bool":
				actual = CONF.GetBool(tt.group, tt.key)
			}
			if !reflect.DeepEqual(tt.expected, actual) {
				t.Fatalf("want value %v(%T) but got %v(%T)", tt.expected, tt.expected, actual, actual)
			}
		})
	}
}
