package main

import (
	"net/http"
	"strings"

	"github.com/andrewmeissner/react-gin-binary/ui"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-gonic/gin"
)

type binaryFileSystem struct {
	fs http.FileSystem
}

func (b *binaryFileSystem) Open(name string) (http.File, error) {
	return b.fs.Open(name)
}

func (b *binaryFileSystem) Exists(prefix, filepath string) bool {
	if p := strings.TrimPrefix(filepath, prefix); len(p) > len(filepath) {
		if _, err := b.fs.Open(p); err != nil {
			return false
		}
		return true
	}
	return false
}

// BinaryFileSystem ...
func BinaryFileSystem(root string) *binaryFileSystem {
	return &binaryFileSystem{
		fs: &assetfs.AssetFS{
			Asset:     ui.Asset,
			AssetDir:  ui.AssetDir,
			AssetInfo: ui.AssetInfo,
			Prefix:    root,
		},
	}
}

func main() {
	router := gin.Default()

	router.StaticFS("/ui", BinaryFileSystem("ui/build"))

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "/ui")
	})

	api := router.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	router.Run(":8080")
}
