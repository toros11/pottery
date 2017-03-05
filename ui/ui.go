package ui

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/higanworks/envmap"
	"github.com/qb0C80aE/clay/extension"
	"net/http"
	"os"
)

func HookSubmodules() {
}

type RouterInitializer struct {
}

func (_ *RouterInitializer) InitializeEarly(r *gin.Engine) error {
	r.Static("ui/files", "ui/files")
	r.LoadHTMLGlob("ui/templates/*.tmpl")
	envMap := envmap.All()
	if endPoint := os.Getenv("ENDPOINT"); endPoint == "" {
		envMap["ENDPOINT"] = fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	}
	ui := r.Group("/ui")
	{
		ui.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.tmpl", gin.H{"env": envMap})
		})
		ui.GET("/network", func(c *gin.Context) {
			c.HTML(http.StatusOK, "network.tmpl", gin.H{"env": envMap})
		})
		ui.GET("/diagram", func(c *gin.Context) {
			c.HTML(http.StatusOK, "diagram.tmpl", gin.H{"env": envMap})
		})
	}
	return nil
}

func (_ *RouterInitializer) InitializeLate(r *gin.Engine) error {
	return nil
}

func init() {
	extension.RegisterRouterInitializer(&RouterInitializer{})
}
