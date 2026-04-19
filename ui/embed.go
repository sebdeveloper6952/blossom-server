//go:build ui

package ui

import (
	"embed"
	"io/fs"
)

//go:embed all:build
var embedded embed.FS

func FS() (fs.FS, error) {
	return fs.Sub(embedded, "build")
}
