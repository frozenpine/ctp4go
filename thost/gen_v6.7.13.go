//go:build v6.7.13

package thost

//go:generate go run ../gen/main.go -dep ../dependencies -plat future -sdk name=mduser,ver=v6.7.13 -output thost

//go:generate gofmt -w ./types
