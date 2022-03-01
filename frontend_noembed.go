//go:build !embed

package main

import "embed"

var embeddedFrontend embed.FS

var defaultFrontend = "frontend"
var frontendHelpExt string
