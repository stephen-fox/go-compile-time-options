// Package compileoptions implements compile time options for the 'myapp'
// command line application using Go's build tags feature.
//
// The following tags are supported:
//	- debug
//	- release
//
// These tags can be leveraged with:
//	go run -tags <tag> cmd/myapp/main.go
package compileoptions
