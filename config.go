package config

import "os"

func LoadConfig(path string) {
	f, err := os.ReadFile(path)
	
}