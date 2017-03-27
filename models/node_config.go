package models

import (
	"github.com/qb0C80aE/clay/extension"
	loamModels "github.com/qb0C80aE/loam/models"
)

type NodeConfig struct {
	ID               int              `json:"id" form:"id" gorm:"primary_key;AUTO_INCREMENT"`
	EnvironmentID    int              `json:"environment_id" gorm:"index" sql:"type:integer references environments(id) on delete cascade"`
	Environment      *Environment     `json:"environment"`
	NodeID           int              `json:"node_id" gorm:"index" sql:"type:integer references nodes(id) on delete cascade"`
	Node             *loamModels.Node `json:"node"`
	FirmwareVersion  string           `json:"firmware_version" gorm:"not null"`
	InitializeConfig string           `json:"initialize_config" gorm:"not null"`
	Config           string           `json:"config" gorm:"not null"`
}

var NodeConfigModel = &NodeConfig{}

func init() {
	extension.RegisterModelType(NodeConfigModel)
}
