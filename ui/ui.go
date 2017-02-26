package ui

import (
	"fmt"
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/higanworks/envmap"
	"github.com/qb0C80aE/clay/extension"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

func HookSubmodules() {
}

type binaryFileSystem struct {
	fs http.FileSystem
}

func (b *binaryFileSystem) Open(name string) (http.File, error) {
	return b.fs.Open(name)
}

func (b *binaryFileSystem) Exists(prefix string, filepath string) bool {

	if p := strings.TrimPrefix(filepath, prefix); len(p) < len(filepath) {
		if _, err := b.fs.Open(p); err != nil {
			return false
		}
		return true
	}
	return false
}

func BinaryFileSystem(root string) *binaryFileSystem {
	fs := &assetfs.AssetFS{Asset, AssetDir, AssetInfo, root}
	return &binaryFileSystem{
		fs,
	}
}

func loadTemplates(filenames ...string) (*template.Template, error) {
	var t *template.Template = nil
	for _, filename := range filenames {
		templateString, _ := Asset(filename)
		name := filepath.Base(filename)
		var tmpl *template.Template
		if t == nil {
			t = template.New(name)
		}
		if name == t.Name() {
			tmpl = t
		} else {
			tmpl = t.New(name)
		}
		_, err := tmpl.Parse(string(templateString))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

type RouterInitializer struct {
}

type PageTemplate struct {
	Name string
}

func (_ *RouterInitializer) InitializeEarly(r *gin.Engine) error {
	r.Use(static.Serve("/ui/files", BinaryFileSystem("ui/files")))

	templateBase := "ui/templates"
	templateFileNames := []string{
		fmt.Sprintf("%s/%s", templateBase, "diagram.tmpl"),
		fmt.Sprintf("%s/%s", templateBase, "dialog.tmpl"),
		fmt.Sprintf("%s/%s", templateBase, "footerpart.tmpl"),
		fmt.Sprintf("%s/%s", templateBase, "headerpart.tmpl"),
		fmt.Sprintf("%s/%s", templateBase, "index.tmpl"),
		fmt.Sprintf("%s/%s", templateBase, "navigation_bar.tmpl"),
		fmt.Sprintf("%s/%s", templateBase, "network.tmpl"),
		fmt.Sprintf("%s/%s", templateBase, "network_edit_node.tmpl"),
		fmt.Sprintf("%s/%s", templateBase, "network_edit_port.tmpl"),
	}
	htmlTemplate, err := loadTemplates(templateFileNames...)
	if err != nil {
		log.Fatal(err)
		return err
	}
	r.SetHTMLTemplate(htmlTemplate)

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
		ui.GET("/environment", func(c *gin.Context) {
			c.HTML(http.StatusOK, "environment.tmpl", gin.H{"env": envMap, "resource": "environment"})
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
