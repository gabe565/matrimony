package datadir

import (
	"io/fs"
	"os"
	"path"
)

func PublicDir() string {
	return path.Join(DataDir, "public")
}

func PublicFS() fs.FS {
	return os.DirFS(PublicDir())
}
