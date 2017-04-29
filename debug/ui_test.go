// +build debug
// exec 'go test -v -tags=debug ./debug/...' and access localhost:8080/ui for GUI debugging or manual testing

package debug

import (
	"fmt"
	_ "github.com/qb0C80aE/clay/controllers" // Install Clay controller sub module by importing
	"github.com/qb0C80aE/clay/db"
	_ "github.com/qb0C80aE/clay/logics" // Install Clay logic sub module by importing
	_ "github.com/qb0C80aE/clay/models" // Install Clay model sub module by importing
	"github.com/qb0C80aE/clay/server"
	_ "github.com/qb0C80aE/loam"    // Install Loam sub module by importing
	_ "github.com/qb0C80aE/pottery" // Install Pottery sub module by importing
	"os"
	"path/filepath"
	"testing"
)

func TestOperateUI(t *testing.T) {
	projectDir, _ := filepath.Abs("..")
	os.Chdir(projectDir)

	host := "localhost"
	port := "8080"

	database := db.Connect()
	s := server.Setup(database)

	s.Run(fmt.Sprintf("%s:%s", host, port))
}
