package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	clayControllers "github.com/qb0C80aE/clay/controllers"
	"github.com/qb0C80aE/clay/extensions"
	"github.com/qb0C80aE/pottery/logics"
	"github.com/qb0C80aE/pottery/models"
)

type logicalDiagramController struct {
	*clayControllers.BaseController
}

func newLogicalDiagramController() extensions.Controller {
	controller := &logicalDiagramController{
		BaseController: clayControllers.NewBaseController(
			models.SharedDiagramModel(),
			logics.UniqueLogicalDiagramLogic(),
		),
	}
	controller.SetOutputter(controller)
	return controller
}

func (controller *logicalDiagramController) RouteMap() map[int]map[string]gin.HandlerFunc {
	url := fmt.Sprintf("%s/logical", controller.ResourceName())
	routeMap := map[int]map[string]gin.HandlerFunc{
		extensions.MethodGet: {
			url: controller.GetSingle,
		},
	}
	return routeMap
}

var uniqueLogicalDiagramController = newLogicalDiagramController()

func init() {
	extensions.RegisterController(uniqueLogicalDiagramController)
}
