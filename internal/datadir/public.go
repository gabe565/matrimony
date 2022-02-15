package datadir

import (
	"io/fs"
	"os"
	"path/filepath"
)

func PublicDir() string {
	return filepath.Join(DataDir, "public")
}

func PublicFS() fs.FS {
	return os.DirFS(PublicDir())
}
