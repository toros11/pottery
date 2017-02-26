package controllers

import (
	"github.com/gin-gonic/gin"
	clayControllers "github.com/qb0C80aE/clay/controllers"
	"github.com/qb0C80aE/clay/extension"
	"github.com/qb0C80aE/pottery/logics"
	"github.com/qb0C80aE/pottery/models"
)

type EnvironmentController struct {
	clayControllers.BaseController
}

func init() {
	extension.RegisterController(NewEnvironmentController())
}

func NewEnvironmentController() *EnvironmentController {
	controller := &EnvironmentController{}
	controller.Initialize()
	return controller
}

func (this *EnvironmentController) Initialize() {
	this.ResourceName = "environment"
	this.Model = models.EnvironmentModel
	this.Logic = logics.EnvironmentLogicInstance
	this.Outputter = this
}

func (this *EnvironmentController) GetRouteMap() map[int]map[string]gin.HandlerFunc {
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
		extension.MethodPatch: {
			resourceSingleUrl: this.Patch,
		},
	}
	return routeMap
}
