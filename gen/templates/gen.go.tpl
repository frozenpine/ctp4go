{{- $sdk := .Sdk -}}
//go:build ignore
package {{ $sdk.Version | ReplaceAll "." "_"}}

//go:generate go run ../../gen/main.go -dep ../../dependencies -plat {{ .Platform }} -sdk name={{ $sdk.Name }},ver={{ $sdk.Version }} -output api

//go:generate go run ../../gen/main.go -dep ../../dependencies -plat {{ .Platform }} -sdk name={{ $sdk.Name }},ver={{ $sdk.Version }} -output spi

//go:generate gofmt -w .

//go:generate go build .
