//go:build embed

package main

//go:generate -command npm npm --prefix frontend
//go:generate npm install
//go:generate npm run build

import "embed"

//go:embed frontend/dist
var embeddedFrontend embed.FS

var (
	defaultFrontend string
	frontendHelpExt = " If left empty, embedded assets are used."
)
