package datadir

import (
	flag "github.com/spf13/pflag"
	"path/filepath"
)

const Default = "data"

var DataDir string

func init() {
	flag.StringVarP(&DataDir, "data", "d", Default, "Data directory")
}

func ReplaceIfInDefault(filename string) string {
	if DataDir != Default && filename[0:len(Default)] == Default {
		return filepath.Join(DataDir, filename[len(Default):])
	}
	return filename
}
