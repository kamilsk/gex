# gex
[![Go project version](https://badge.fury.io/go/github.com%2Fizumin5210%2Fgex.svg)](https://badge.fury.io/go/github.com%2Fizumin5210%2Fgex)
[![Go Report Card](https://goreportcard.com/badge/github.com/izumin5210/gex)](https://goreportcard.com/report/github.com/izumin5210/gex)
[![License](https://img.shields.io/github/license/izumin5210/gex.svg)](./LICENSE)

The implementation of clarify best practice for tool dependencies.

See https://github.com/golang/go/issues/25922#issuecomment-412992431


## Installation
To install gex, you can use `go get`.

```
$ go get github.com/izumin5210/gex/cmd/gex
```


## Usage

### `gex --add [packages...]`
Add a new tool to dependencies:

```
$ gex --add github.com/golang/mock/mockgen
```

The tool will be managed in `tools.go` and its version will be managed by [Go Modules](https://github.com/golang/go/wiki/Modules).

```
$ cat tools.go
// Code generated by github.com/izumin5210/gex. DO NOT EDIT.

// +build tools

package tools

// tool dependencies
import (
        _ "github.com/golang/mock/mockgen"
)

$ cat go.mod | grep mock
        github.com/golang/mock v1.1.1 // indirect
```


### `gex [command] [args...]`
Execute command that managed in `tools.go` and `go.mod`.
`gex` will build the executable binary automatically if needed.

```
$ gex mockgen
# prints mockgen's help text...
```