package config

type IniFile struct {
	sections []Section
}

func NewIniFile() *IniFile {
	return &IniFile{
		sections: make([]Section, 0),
	}
}

func LoadIniFile(filePath string) *IniFile {
	// TODO: optimize

}
