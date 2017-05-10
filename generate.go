// +build generate

package pottery

//go:generate go-bindata -pkg ui -o ui/bindata.go ui/files/... ui/templates/...
