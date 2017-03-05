package controllers

import (
	"github.com/gin-gonic/gin"
	clayControllers "github.com/qb0C80aE/clay/controllers"
	"github.com/qb0C80aE/clay/extension"
	"github.com/qb0C80aE/pottery/logics"
	"github.com/qb0C80aE/pottery/models"
)

type PhysicalDiagramController struct {
	clayControllers.BaseController
}

type LogicalDiagramController struct {
	clayControllers.BaseController
}

func init() {
	extension.RegisterController(NewPhysicalDiagramController())
	extension.RegisterController(NewLogicalDiagramController())
}

func NewPhysicalDiagramController() *PhysicalDiagramController {
	controller := &PhysicalDiagramController{}
	controller.Initialize()
	return controller
}

func NewLogicalDiagramController() *LogicalDiagramController {
	controller := &LogicalDiagramController{}
	controller.Initialize()
	return controller
}

func (this *PhysicalDiagramController) Initialize() {
	this.ResourceName = "diagram_physical"
	this.Model = models.DiagramModel
	this.Logic = logics.PhysicalDiagramLogicInstance
	this.Outputter = this
}

func (this *LogicalDiagramController) Initialize() {
	this.ResourceName = "diagram_logical"
	this.Model = models.DiagramModel
	this.Logic = logics.LogicalDiagramLogicInstance
	this.Outputter = this
}

func (this *PhysicalDiagramController) GetRouteMap() map[int]map[string]gin.HandlerFunc {
	url := "diagrams/physical"
	routeMap := map[int]map[string]gin.HandlerFunc{
		extension.MethodGet: {
			url: this.GetSingle,
		},
	}
	return routeMap
}

func (this *LogicalDiagramController) GetRouteMap() map[int]map[string]gin.HandlerFunc {
	url := "diagrams/logical"
	routeMap := map[int]map[string]gin.HandlerFunc{
		extension.MethodGet: {
			url: this.GetSingle,
		},
	}
	return routeMap
}
