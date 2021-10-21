package static

import "embed"

//go:embed admin/* web/*
var StaticDir embed.FS
