package assets

import "embed"

//go:embed "css/*" "js/*"
var Files embed.FS
