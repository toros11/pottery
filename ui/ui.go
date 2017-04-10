package ui

import (
	"github.com/gin-gonic/gin"
	"github.com/higanworks/envmap"
	"github.com/qb0C80aE/clay/extensions"
	"net/http"
)

type RouterInitializer struct {
}

type PageTemplate struct {
	Name string
}

func (routerInitializer *RouterInitializer) InitializeEarly(r *gin.Engine) error {
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
		ui.GET("/diagram", func(c *gin.Context) {
			c.HTML(http.StatusOK, "diagram.tmpl", gin.H{"env": envMap, "category": "design"})
		})
	}
	return nil
}

func (routerInitializer *RouterInitializer) InitializeLate(r *gin.Engine) error {
	return nil
}

func init() {
	extensions.RegisterRouterInitializer(&RouterInitializer{})
}
