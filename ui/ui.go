package ui

import (
	"github.com/gin-gonic/gin"
	"github.com/higanworks/envmap"
	"github.com/qb0C80aE/clay/extension"
	"net/http"
)

func HookSubmodules() {
}

type RouterInitializer struct {
}

type PageTemplate struct {
	Name string
}

func (_ *RouterInitializer) InitializeEarly(r *gin.Engine) error {
	r.Static("ui/files", "ui/files")
	r.LoadHTMLGlob("ui/templates/*.tmpl")
	envMap := envmap.All()
	ui := r.Group("/ui")
	{
		ui.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.tmpl", gin.H{"env": envMap, "resource": "index"})
		})
		ui.GET("/network", func(c *gin.Context) {
			c.HTML(http.StatusOK, "network.tmpl", gin.H{"env": envMap, "resource": "network"})
		})
		ui.GET("/diagram", func(c *gin.Context) {
			c.HTML(http.StatusOK, "diagram.tmpl", gin.H{"env": envMap, "resource": "diagram"})
		})
		ui.GET("/requirement", func(c *gin.Context) {
			c.HTML(http.StatusOK, "requirement.tmpl", gin.H{"env": envMap, "resource": "requirement"})
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
