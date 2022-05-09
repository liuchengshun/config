package config

import "testing"

func TestTrim(t *testing.T) {
	tests := []struct {
		name   string
		text   string
		expect string
	}{
		{"trim[1]", "#server", ""},
		{"trim[2]", "port=8989 #port", "port=8989"},
		{"trim[3]", "server=127.0.0.1    #server", "server=127.0.0.1"},
		{"trim[4]", "testing    #server", "testing"},
		{"trim[4]", "    testing    #server", "testing"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			confline := newConfLine(tt.text)
			actual := confline.trim()
			if actual != tt.expect {
				t.Fatalf("want %s but got %s", tt.expect, actual)
			}
		})
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		name      string
		text      string
		lineT     int
		key       string
		value     string
		groupName string
	}{
		{"parse[0]", "#system", lineEmpty, "", "", ""},
		{"parse[1]", "port = 8989 #system", lineKeyValue, "port", "8989", ""},
		// {"parse[0]", "", 0, "", "", ""},
		// {"parse[0]", "", 0, "", "", ""},
		// {"parse[0]", "", 0, "", "", ""},
		// {"parse[0]", "", 0, "", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			confline := newConfLine(tt.text)
			lineT, err := confline.parse()
			if err != nil {
				t.Fatalf("want nil bot got error %v", err)
				return
			}

			if lineT != tt.lineT {
				t.Fatalf("want line type %d but got %d", tt.lineT, lineT)
			}

			key, value := confline.keyValue()
			if key != tt.key {
				t.Fatalf("want key %s but got %s", tt.key, key)
			}
			if value != tt.value {
				t.Fatalf("want value %s but got %s", tt.value, value)
			}

			groupName := confline.confGroupName()
			if groupName != tt.groupName {
				t.Fatalf("want group name %s bot got %s", tt.groupName, groupName)
			}
		})
	}
}
