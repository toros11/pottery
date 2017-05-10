// +build debug
// exec 'go test -v -tags=debug ./debug/...' and access localhost:8080/ui for GUI debugging or manual testing

package debug

import (
	"github.com/qb0C80aE/clay/extensions"
	_ "github.com/qb0C80aE/clay/runtime" // Install Runtime sub module by importing
	_ "github.com/qb0C80aE/loam"         // Install Loam sub module by importing
	_ "github.com/qb0C80aE/pottery"      // Install Pottery sub module by importing
	"os"
	"path/filepath"
	"testing"
)

func TestOperateUI(t *testing.T) {
	projectDir, _ := filepath.Abs("..")
	os.Chdir(projectDir)

	extensions.RegisteredRuntime().Run()
}
