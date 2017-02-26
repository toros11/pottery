package models

import (
	"github.com/qb0C80aE/clay/extension"
	clayModels "github.com/qb0C80aE/clay/models"
)

type Environment struct {
	ID                    int                  `json:"id" form:"id" gorm:"primary_key;AUTO_INCREMENT"`
	TemplateID            int                  `json:"template_id" gorm:"index" sql:"type:integer references templates(id) on delete set null"`
	Template              *clayModels.Template `json:"template"`
	TestCaseID            int                  `json:"test_case_id" gorm:"index" sql:"type:integer references test_cases(id) on delete set null"`
	TestCase              *TestCase            `json:"test_case"`
	GitRepositoryURI      string               `json:"git_repository_uri" gorm:"not null"`
	GitUserName           string               `json:"git_user_name" gorm:"not null"`
	GitUserEmail          string               `json:"git_user_email" gorm:"not null"`
	DesignFileName        string               `json:"design_file_name" gorm:"not null"`
	TemplateFileName      string               `json:"template_file_name" gorm:"not null"`
	TestCaseDirectoryName string               `json:"test_case_directory_name" gorm:"not null"`
}

var EnvironmentModel = &Environment{}

func init() {
	extension.RegisterModelType(EnvironmentModel)
}
