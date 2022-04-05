# Collection

![Build Status](https://github.com/wilhelm-murdoch/go-collection/actions/workflows/go.yml/badge.svg)
[![GoDoc](https://godoc.org/github.com/wilhelm-murdoch/go-collection?status.svg)](https://pkg.go.dev/github.com/wilhelm-murdoch/go-collection)
[![Go report](https://goreportcard.com/badge/github.com/wilhelm-murdoch/go-collection)](https://goreportcard.com/report/github.com/wilhelm-murdoch/go-collection)

A generic collection for Go with a few convenient methods.
# Install
```
go get github.com/wilhelm-murdoch/go-collection
```
{{ range . }}{{ range .Files.Items }}{{ $path := .Path }}{{ range .Functions.Items }}{{ if and (ne .Name "main") (not .IsTest) (not .IsExample) (not .IsBenchmark)}}### Func {{ .Name }}
* `{{ trim .Signature }}` [#]()
* `{{ $path }}:{{ .LineStart }}-{{ .LineEnd }}` [#]()

{{ .Comment | replace "\n" "" }}
{{ range .Examples }}
```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
{{ indent 4 .Body }}
}
```
{{ if .Output }}```go
{{ .Output }}
```{{ end }}{{ end }}
{{ end }}{{ end }}{{ end }}{{ end }}
# License
Copyright © {{ now | date "2006" }} [Wilhelm Murdoch](https://wilhelm.codes).

This project is [MIT](./LICENSE) licensed.
