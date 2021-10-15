package static

import "embed"

//go:embed admin
var AdminStaticDir embed.FS

//go:embed web
var WebStaticDir embed.FS
