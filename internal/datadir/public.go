package datadir

import (
	"io/fs"
	"os"
	"path/filepath"
)

const PublicDir = "public"

func PublicFS() fs.FS {
	return os.DirFS(filepath.Join(DataDir, "public"))
}
