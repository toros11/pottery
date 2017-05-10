// +build debug
// exec 'go generate -tags=debug ./debug/...' to generate bindata for debugging

package debug

//go:generate go-bindata -debug=true -pkg ui -o ../ui/bindata.go -prefix=../ ../ui/files/... ../ui/templates/...
