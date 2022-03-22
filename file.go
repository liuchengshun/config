package config

import "path/filepath"

type fileType string

const (
	fileTypeJSON   fileType = "config_type_json"
	fileTypeKeyVal fileType = "config_type_key_value"

	jsonExtName   = ".json"
	keyValExtName = ".conf"
)

func parseFileType(filePath string) fileType {
	switch filepath.Ext(filePath) {
	case jsonExtName:
		return fileTypeJSON
	case keyValExtName:
		return fileTypeKeyVal
	default:
		panic("only support json or kay/value configuration, the suffix of the former is JSON, the latter is conf")
	}
}
