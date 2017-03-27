package logics

import (
	"github.com/jinzhu/gorm"
	"github.com/qb0C80aE/clay/extension"
	clayModels "github.com/qb0C80aE/clay/models"
	"github.com/qb0C80aE/clay/utils/mapstruct"
	"github.com/qb0C80aE/pottery/models"
	"strconv"
)

type NodeConfigLogic struct {
}

func (_ *NodeConfigLogic) GetSingle(db *gorm.DB, id string, queryFields string) (interface{}, error) {

	nodeConfig := &models.NodeConfig{}

	if err := db.Select(queryFields).First(nodeConfig, id).Error; err != nil {
		return nil, err
	}

	return nodeConfig, nil

}

func (_ *NodeConfigLogic) GetMulti(db *gorm.DB, queryFields string) ([]interface{}, error) {

	nodeConfigs := []*models.NodeConfig{}

	if err := db.Select(queryFields).Find(&nodeConfigs).Error; err != nil {
		return nil, err
	}

	result := make([]interface{}, len(nodeConfigs))
	for i, data := range nodeConfigs {
		result[i] = data
	}

	return result, nil

}

func (_ *NodeConfigLogic) Create(db *gorm.DB, data interface{}) (interface{}, error) {

	nodeConfig := data.(*models.NodeConfig)

	if err := db.Create(nodeConfig).Error; err != nil {
		return nil, err
	}

	return nodeConfig, nil

}

func (_ *NodeConfigLogic) Update(db *gorm.DB, id string, data interface{}) (interface{}, error) {

	nodeConfig := data.(*models.NodeConfig)
	nodeConfig.ID, _ = strconv.Atoi(id)

	if err := db.Save(&nodeConfig).Error; err != nil {
		return nil, err
	}

	return nodeConfig, nil

}

func (_ *NodeConfigLogic) Delete(db *gorm.DB, id string) error {

	nodeConfig := &models.NodeConfig{}

	if err := db.First(&nodeConfig, id).Error; err != nil {
		return err
	}

	if err := db.Delete(&nodeConfig).Error; err != nil {
		return err
	}

	return nil

}

func (this *NodeConfigLogic) Patch(db *gorm.DB, id string, _ string) (interface{}, error) {
	return nil, nil
}

func (_ *NodeConfigLogic) Options(_ *gorm.DB) error {
	return nil
}

func (_ *NodeConfigLogic) ExtractFromDesign(db *gorm.DB) (string, interface{}, error) {
	nodeConfigs := []*models.NodeConfig{}
	if err := db.Select("*").Find(&nodeConfigs).Error; err != nil {
		return "", nil, err
	}
	return "node_configs", nodeConfigs, nil
}

func (_ *NodeConfigLogic) DeleteFromDesign(db *gorm.DB) error {
	return db.Exec("delete from node_configs;").Error
}

func (_ *NodeConfigLogic) LoadToDesign(db *gorm.DB, data interface{}) error {
	container := []*models.NodeConfig{}
	design := data.(*clayModels.Design)
	if value, exists := design.Content["node_configs"]; exists {
		if err := mapstruct.MapToStruct(value.([]interface{}), &container); err != nil {
			return err
		}
		for _, nodeConfig := range container {
			nodeConfig.Environment = nil
			nodeConfig.Node = nil
			if err := db.Create(nodeConfig).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

var NodeConfigLogicInstance = &NodeConfigLogic{}

func init() {
	extension.RegisterDesignAccessor(NodeConfigLogicInstance)
}
