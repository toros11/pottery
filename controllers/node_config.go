package controllers

import (
	"github.com/gin-gonic/gin"
	clayControllers "github.com/qb0C80aE/clay/controllers"
	"github.com/qb0C80aE/clay/extension"
	"github.com/qb0C80aE/pottery/logics"
	"github.com/qb0C80aE/pottery/models"
)

type NodeConfigController struct {
	clayControllers.BaseController
}

func init() {
	extension.RegisterController(NewNodeConfigController())
}

func NewNodeConfigController() *NodeConfigController {
	controller := &NodeConfigController{}
	controller.Initialize()
	return controller
}

func (this *NodeConfigController) Initialize() {
	this.ResourceName = "node_config"
	this.Model = models.NodeConfigModel
	this.Logic = logics.NodeConfigLogicInstance
	this.Outputter = this
}

func (this *NodeConfigController) GetRouteMap() map[int]map[string]gin.HandlerFunc {
	resourceSingleUrl := extension.GetResourceSingleUrl(this.ResourceName)
	resourceMultiUrl := extension.GetResourceMultiUrl(this.ResourceName)

	routeMap := map[int]map[string]gin.HandlerFunc{
		extension.MethodGet: {
			resourceSingleUrl: this.GetSingle,
			resourceMultiUrl:  this.GetMulti,
		},
		extension.MethodPost: {
			resourceMultiUrl: this.Create,
		},
		extension.MethodPut: {
			resourceSingleUrl: this.Update,
		},
		extension.MethodDelete: {
			resourceSingleUrl: this.Delete,
		},
	}
	return routeMap
}
