# Collection

![Build Status](https://github.com/wilhelm-murdoch/go-collection/actions/workflows/go.yml/badge.svg)
[![GoDoc](https://godoc.org/github.com/wilhelm-murdoch/go-collection?status.svg)](https://pkg.go.dev/github.com/wilhelm-murdoch/go-collection)
[![Go report](https://goreportcard.com/badge/github.com/wilhelm-murdoch/go-collection)](https://goreportcard.com/report/github.com/wilhelm-murdoch/go-collection)
[![Stability: Active](https://masterminds.github.io/stability/active.svg)](https://masterminds.github.io/stability/active.html)

A generic collection for Go with a few convenient methods. 

There are far more comprehensive modules out there, but this one works quite well for my purposes.
# Install
```
go get github.com/wilhelm-murdoch/go-collection
```
# Reference

{{ range . }}{{ range .Files.Items }}{{ range .Functions.Items }}{{ if and (.IsExported) (ne .Name "main") (not .IsTest) (not .IsExample) (not .IsBenchmark)}}* [{{ .Name }}](#Function-{{ .Name }})
{{ end }}{{ end }}{{ end }}{{ end }}

{{ range . }}{{ range .Files.Items }}{{ $path := .Path }}{{ range .Functions.Items }}{{ if and (.IsExported) (ne .Name "main") (not .IsTest) (not .IsExample) (not .IsBenchmark)}}### Function `{{ .Name }}`
* `{{ trim .Signature }}` [#]({{ $path }}#L{{ .LineStart }})
* `{{ $path }}:{{ .LineStart }}:{{ .LineEnd }}` [#]({{ $path }}#L{{ .LineStart }}-L{{ .LineEnd }})

{{ .Doc | replace "\n" "" }}
{{ range .Examples.Items }}
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