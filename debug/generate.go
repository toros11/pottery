// +build debug
// execute 'go generate -tags=debug debug/generate.go' or 'go generate -tags=debug ./...' to generate bindata.go for debugging manually

package main

//go:generate go-bindata -debug=true -pkg ui -o ../ui/bindata.go -prefix=../ ../ui/files/... ../ui/templates/...
