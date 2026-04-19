//go:build !ui

package ui

import (
	"embed"
	"io/fs"
)

//go:embed stub
var embedded embed.FS

func FS() (fs.FS, error) {
	return fs.Sub(embedded, "stub")
}
