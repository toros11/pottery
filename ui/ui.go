package ui

import (
	"github.com/gin-gonic/gin"
	"github.com/higanworks/envmap"
	"github.com/qb0C80aE/clay/extensions"
	"net/http"
)

type routerInitializer struct {
}

func newRouterInitializer() *routerInitializer {
	return &routerInitializer{}
}

func (routerInitializer *routerInitializer) InitializeEarly(r *gin.Engine) error {
	r.Static("ui/files", "ui/files")
	r.LoadHTMLGlob("ui/templates/*.tmpl")
	envMap := envmap.All()
	ui := r.Group("/ui")
	{
		ui.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.tmpl", gin.H{"env": envMap, "category": "home"})
		})
		ui.GET("/network", func(c *gin.Context) {
			c.HTML(http.StatusOK, "network.tmpl", gin.H{"env": envMap, "category": "design"})
		})
		ui.GET("/design", func(c *gin.Context) {
			c.HTML(http.StatusOK, "design.tmpl", gin.H{"env": envMap, "category": "design"})
		})
		ui.GET("/diagram", func(c *gin.Context) {
			c.HTML(http.StatusOK, "diagram.tmpl", gin.H{"env": envMap, "category": "design"})
		})
		ui.GET("/template", func(c *gin.Context) {
			c.HTML(http.StatusOK, "template.tmpl", gin.H{"env": envMap, "category": "process"})
		})
	}
	return nil
}

func (routerInitializer *routerInitializer) InitializeLate(r *gin.Engine) error {
	return nil
}

var uniqueRouterInitializer = newRouterInitializer()

func init() {
	extensions.RegisterRouterInitializer(uniqueRouterInitializer)
}
