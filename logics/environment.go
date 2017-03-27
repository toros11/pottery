package logics

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/qb0C80aE/clay/extension"
	clayLogics "github.com/qb0C80aE/clay/logics"
	clayModels "github.com/qb0C80aE/clay/models"
	"github.com/qb0C80aE/clay/utils/mapstruct"
	"github.com/qb0C80aE/pottery/models"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	tplpkg "text/template"
	"time"
)

type EnvironmentLogic struct {
}

func initGitRepository(environment *models.Environment) error {
	cmd := exec.Command("mkdir", "-p", environment.GitRepositoryURI)
	cmdMessage, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(cmdMessage))
	}

	cmd = exec.Command("git", "status")
	cmd.Dir = environment.GitRepositoryURI
	if err := cmd.Run(); err != nil {
		cmd := exec.Command("git", "init")
		cmd.Dir = environment.GitRepositoryURI
		cmd.Run()
	}

	cmd = exec.Command("git", "config", "--local", "user.name", environment.GitUserName)
	cmd.Dir = environment.GitRepositoryURI
	cmdMessage, err = cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(cmdMessage))
	}

	cmd = exec.Command("git", "config", "--local", "user.email", environment.GitUserEmail)
	cmd.Dir = environment.GitRepositoryURI
	cmdMessage, err = cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(cmdMessage))
	}

	return nil
}

func updateDesignFile(db *gorm.DB, environment *models.Environment) error {
	design, err := clayLogics.DesignLogicInstance.GetSingle(db, "", "*")

	if err != nil {
		return err
	}

	jsonString, err := json.MarshalIndent(design, "", "    ")
	if ioutil.WriteFile(fmt.Sprintf("%s/%s", environment.GitRepositoryURI, environment.DesignFileName), jsonString, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func updateTemplateFile(db *gorm.DB, environment *models.Environment) error {
	templateID := fmt.Sprintf("%d", environment.TemplateID)
	template, err := clayLogics.TemplateLogicInstance.Patch(db, templateID, "*")

	if err != nil {
		return err
	}

	if ioutil.WriteFile(fmt.Sprintf("%s/%s", environment.GitRepositoryURI, environment.TemplateFileName), ([]byte)(template.(string)), os.ModePerm); err != nil {
		return err
	}

	return nil
}

func updateTestCaseFile(db *gorm.DB, environment *models.Environment) error {
	testCaseID := fmt.Sprintf("%d", environment.TestCaseID)
	testRunnerScriptTemplate, testCommands, err := generateTestScripts(db, testCaseID)

	if err != nil {
		return err
	}

	cmd := exec.Command("rm", "-rf", environment.TestCaseDirectoryName)
	cmd.Dir = environment.GitRepositoryURI
	cmdMessage, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(cmdMessage))
	}

	cmd = exec.Command("mkdir", environment.TestCaseDirectoryName)
	cmd.Dir = environment.GitRepositoryURI
	cmdMessage, err = cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(cmdMessage))
	}

	var docServerScript bytes.Buffer
	tplTestRunnerScript, _ := tplpkg.New("template_test_runner_script").Parse(testRunnerScriptTemplate)
	tplTestRunnerScript.Execute(&docServerScript, environment)
	testRunnerScript := docServerScript.String()

	if ioutil.WriteFile(fmt.Sprintf("%s/test_runner.sh",
		environment.GitRepositoryURI),
		([]byte)(testRunnerScript),
		os.ModePerm); err != nil {
		return err
	}

	for _, testCommand := range testCommands {
		if ioutil.WriteFile(fmt.Sprintf("%s/%s/%s_server.sh",
			environment.GitRepositoryURI,
			environment.TestCaseDirectoryName,
			testCommand.ServiceName),
			([]byte)(testCommand.SerevrScriptTemplate),
			os.ModePerm); err != nil {
			return err
		}
		if ioutil.WriteFile(fmt.Sprintf("%s/%s/%s_client.sh",
			environment.GitRepositoryURI,
			environment.TestCaseDirectoryName,
			testCommand.ServiceName),
			([]byte)(testCommand.ClientScriptTemplate),
			os.ModePerm); err != nil {
			return err
		}
	}

	return nil
}

func updateNodeConfigFiles(_ *gorm.DB, environment *models.Environment) error {
	for _, config := range environment.NodeConfigs {
		cmd := exec.Command("mkdir", "-p", fmt.Sprintf("%s/%s/%s", environment.GitRepositoryURI, "config", config.Node.Name))
		cmdMessage, err := cmd.CombinedOutput()
		if err != nil {
			return errors.New(string(cmdMessage))
		}
		if ioutil.WriteFile(fmt.Sprintf("%s/%s/%s/initialize.txt",
			environment.GitRepositoryURI,
			"config",
			config.Node.Name),
			([]byte)(config.InitializeConfig),
			os.ModePerm); err != nil {
			return err
		}
		if ioutil.WriteFile(fmt.Sprintf("%s/%s/%s/config.txt",
			environment.GitRepositoryURI,
			"config",
			config.Node.Name),
			([]byte)(config.Config),
			os.ModePerm); err != nil {
			return err
		}
		if ioutil.WriteFile(fmt.Sprintf("%s/%s/%s/firmware_version.txt",
			environment.GitRepositoryURI,
			"config",
			config.Node.Name),
			([]byte)(config.FirmwareVersion),
			os.ModePerm); err != nil {
			return err
		}
	}

	return nil
}

func commit(environment *models.Environment, message string) error {
	cmd := exec.Command("git", "add", ".")
	cmd.Dir = environment.GitRepositoryURI
	cmdMessage, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(cmdMessage))
	}
	cmd = exec.Command("git", "commit", "-m", message)
	cmd.Dir = environment.GitRepositoryURI
	cmdMessage, err = cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(cmdMessage))
	}
	return nil
}

func (_ *EnvironmentLogic) GetSingle(db *gorm.DB, id string, queryFields string) (interface{}, error) {

	environment := &models.Environment{}

	if err := db.Select(queryFields).First(environment, id).Error; err != nil {
		return nil, err
	}

	return environment, nil

}

func (_ *EnvironmentLogic) GetMulti(db *gorm.DB, queryFields string) ([]interface{}, error) {

	environments := []*models.Environment{}

	if err := db.Select(queryFields).Find(&environments).Error; err != nil {
		return nil, err
	}

	result := make([]interface{}, len(environments))
	for i, data := range environments {
		result[i] = data
	}

	return result, nil

}

func (_ *EnvironmentLogic) Create(db *gorm.DB, data interface{}) (interface{}, error) {

	environment := data.(*models.Environment)

	if err := db.Create(environment).Error; err != nil {
		return nil, err
	}

	return environment, nil

}

func (_ *EnvironmentLogic) Update(db *gorm.DB, id string, data interface{}) (interface{}, error) {

	environment := data.(*models.Environment)
	environment.ID, _ = strconv.Atoi(id)

	if err := db.Save(&environment).Error; err != nil {
		return nil, err
	}

	return environment, nil

}

func (_ *EnvironmentLogic) Delete(db *gorm.DB, id string) error {

	environment := &models.Environment{}

	if err := db.First(&environment, id).Error; err != nil {
		return err
	}

	if err := db.Delete(&environment).Error; err != nil {
		return err
	}

	return nil

}

func (this *EnvironmentLogic) Patch(db *gorm.DB, id string, _ string) (interface{}, error) {
	environment := &models.Environment{}

	if err := db.Preload("NodeConfigs").Preload("NodeConfigs.Node").Select("*").First(environment, id).Error; err != nil {
		return nil, err
	}

	if err := initGitRepository(environment); err != nil {
		return "", err
	}

	message := fmt.Sprintf("Automatic commit at %s", time.Now().String())
	if err := updateDesignFile(db, environment); err != nil {
		return "", err
	}
	if err := updateTemplateFile(db, environment); err != nil {
		return "", err
	}
	if err := updateTestCaseFile(db, environment); err != nil {
		return "", err
	}
	if err := updateNodeConfigFiles(db, environment); err != nil {
		return "", err
	}
	if err := commit(environment, message); err != nil {
		return "", err
	}
	return "", nil
}

func (_ *EnvironmentLogic) Options(_ *gorm.DB) error {
	return nil
}

func (_ *EnvironmentLogic) ExtractFromDesign(db *gorm.DB) (string, interface{}, error) {
	environments := []*models.Environment{}
	if err := db.Select("*").Find(&environments).Error; err != nil {
		return "", nil, err
	}
	return "environments", environments, nil
}

func (_ *EnvironmentLogic) DeleteFromDesign(db *gorm.DB) error {
	return db.Exec("delete from environments;").Error
}

func (_ *EnvironmentLogic) LoadToDesign(db *gorm.DB, data interface{}) error {
	container := []*models.Environment{}
	design := data.(*clayModels.Design)
	if value, exists := design.Content["environments"]; exists {
		if err := mapstruct.MapToStruct(value.([]interface{}), &container); err != nil {
			return err
		}
		for _, environment := range container {
			environment.Template = nil
			environment.TestCase = nil
			if err := db.Create(environment).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

var EnvironmentLogicInstance = &EnvironmentLogic{}

func init() {
	extension.RegisterDesignAccessor(EnvironmentLogicInstance)
}
