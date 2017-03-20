package integration

import (
	"github.com/qb0C80aE/loam"
	"github.com/qb0C80aE/pottery"
	"os"
	"testing"
)

// +build integration

func TestMain(m *testing.M) {
	loam.HookSubmodules()
	pottery.HookSubmodules()
	code := m.Run()
	defer os.Exit(code)
}
