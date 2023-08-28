package test

import (
	"os"
	"reflect"
	"testing"

	"github.com/quixote-liu/config"
)

var CONF = config.CONF()

func TestMain(m *testing.M) {
	if err := CONF.LoadConfiguration("./testdata/config.conf"); err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func TestCONFGetValue(t *testing.T) {
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
		{"company group", "int", "company", "age", 85},
		{"company group", "string", "company", "business", "write_code"},
		{"company group", "string", "company", "usedlang", "golang"},
		{"company group", "string", "company", "address", "china,hubei"},
		{"company group", "string", "company", "address_copy", "china_hubei"},
		{"language group", "string", "language", "name", "golang"},
		{"language group", "string", "language", "maxcot", "mouse"},
		{"language group", "string", "language", "version", "1.17.7"},
		{"language group", "bool", "language", "production", true},
		{"redis group", "string", "cache", "url", "redis://127.0.0.1:6379?db=15"},
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

func TestRegisterGroup(t *testing.T) {
	companyGroup := config.NewSection("company")
	// want to change company name.
	companyGroup.SetString("name", "kaopu")
	// add new key/value.
	companyGroup.SetString("address", "china")
	companyGroup.SetBool("open", true)
	// register
	CONF.RegisterGroup(companyGroup)

	cupGroup := config.NewSection("cup")
	cupGroup.SetString("hight", "2dm")
	cupGroup.SetBool("used", true)
	cupGroup.SetInt("age", 2)
	CONF.RegisterGroup(cupGroup)

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
		{"company group", "int", "company", "age", 85},
		{"company group", "string", "company", "business", "write_code"},
		{"company group", "string", "company", "usedlang", "golang"},
		{"language group", "string", "language", "name", "golang"},
		{"language group", "string", "language", "maxcot", "mouse"},
		{"language group", "string", "language", "version", "1.17.7"},
		{"language group", "bool", "language", "production", true},

		// test new group key value.
		{"company group", "string", "company", "address", "china,hubei"},
		{"company group", "bool", "company", "open", true},
		{"cup group", "string", "cup", "hight", "2dm"},
		{"cup group", "bool", "cup", "used", true},
		{"cup group", "int", "cup", "age", 2},
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
