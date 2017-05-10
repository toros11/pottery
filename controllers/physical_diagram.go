package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	clayControllers "github.com/qb0C80aE/clay/controllers"
	"github.com/qb0C80aE/clay/extensions"
	"github.com/qb0C80aE/pottery/logics"
	"github.com/qb0C80aE/pottery/models"
)

type physicalDiagramController struct {
	*clayControllers.BaseController
}

func newPhysicalDiagramController() extensions.Controller {
	controller := &physicalDiagramController{
		BaseController: clayControllers.NewBaseController(
			models.SharedDiagramModel(),
			logics.UniquePhysicalDiagramLogic(),
		),
	}
	controller.SetOutputter(controller)
	return controller
}

func (controller *physicalDiagramController) RouteMap() map[int]map[string]gin.HandlerFunc {
	url := fmt.Sprintf("%s/physical", controller.ResourceName())
	routeMap := map[int]map[string]gin.HandlerFunc{
		extensions.MethodGet: {
			url: controller.GetSingle,
		},
	}
	return routeMap
}

var uniquePhysicalDiagramController = newPhysicalDiagramController()

func init() {
	extensions.RegisterController(uniquePhysicalDiagramController)
}
