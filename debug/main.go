// +build debug
// execute 'go run -tags=debug ./debug/main.go' and access localhost:8080/ui for GUI debugging or manual testing

package main

import (
	"github.com/qb0C80aE/clay/extensions"
	_ "github.com/qb0C80aE/clay/runtime" // Import runtime package to register Clay runtime
	_ "github.com/qb0C80aE/loam"         // Install Loam module by importing
	_ "github.com/qb0C80aE/pottery"      // Install Pottery module by importing
)

func main() {
	extensions.RegisteredRuntime().Run()
}
