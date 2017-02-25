package pottery

import (
	"github.com/qb0C80aE/pottery/controllers"
	"github.com/qb0C80aE/pottery/logics"
	"github.com/qb0C80aE/pottery/models"
	"github.com/qb0C80aE/pottery/ui"
)

func HookSubmodules() {
	controllers.HookSubmodules()
	logics.HookSubmodules()
	models.HookSubmodules()
	ui.HookSubmodules()
}
