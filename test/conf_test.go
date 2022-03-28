package test

import (
	"config"
	"testing"
)

func TestLoadConfiguration(t *testing.T) {
	err := config.LoadConfiguration("./testdata/config.conf")
	if err != nil {
		t.Fatal(err)
	}

	CONF := config.CONF

	if CONF.GetString("server", "port") != "8080" {
		t.Fatalf("want server.port 8080 but got %s", CONF.GetString("server", "port"))
	}
}
