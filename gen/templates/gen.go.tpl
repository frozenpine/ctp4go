{{- $sdk := .Sdk -}}
package 

//go:generate go run ../../gen/main.go -dep ../../dependencies -plat {{ .Platform }} -sdk name={{ $sdk.Name }},ver={{ $sdk.Version }} -out api

//go:generate go run ../../gen/main.go -dep ../../dependencies -plat {{ .Platform }} -sdk name={{ $sdk.Name }},ver={{ $sdk.Version }} -out spi

//go:generate gofmt -w .

//go:generate go build .
