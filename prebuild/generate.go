// +build prebuild

package prebuild

//go:generate go-bindata -pkg ui -prefix ../ -o ../ui/bindata.go ../ui/files/... ../ui/templates/...
