package httpapi

import (
	"io/fs"
	"mime"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func serveUI(uiFS fs.FS) gin.HandlerFunc {
	return func(c *gin.Context) {
		p := strings.TrimPrefix(c.Param("filepath"), "/")
		if p == "" {
			p = "index.html"
		}

		data, err := fs.ReadFile(uiFS, p)
		if err != nil {
			data, err = fs.ReadFile(uiFS, "index.html")
			if err != nil {
				c.Status(http.StatusNotFound)
				return
			}
			p = "index.html"
		}

		ct := mime.TypeByExtension(filepath.Ext(p))
		if ct == "" {
			ct = http.DetectContentType(data)
		}
		c.Data(http.StatusOK, ct, data)
	}
}
