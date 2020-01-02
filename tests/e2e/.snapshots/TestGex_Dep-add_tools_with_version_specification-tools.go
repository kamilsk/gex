// Code generated by github.com/kamilsk/egg. DO NOT EDIT.

// +build tools

package tools

import (
	_ "github.com/gogo/protobuf/protoc-gen-gogo"
	_ "github.com/gogo/protobuf/protoc-gen-gogofast"
	_ "github.com/golang/mock/mockgen"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger"
	_ "github.com/izumin5210/gex/cmd/gex"
	_ "golang.org/x/lint/golint"
)

//go:generate go build -v -o=${ROOT}bin/protoc-gen-gogo ./vendor/github.com/gogo/protobuf/protoc-gen-gogo
//go:generate go build -v -o=${ROOT}bin/protoc-gen-gogofast ./vendor/github.com/gogo/protobuf/protoc-gen-gogofast
//go:generate go build -v -o=${ROOT}bin/mockgen ./vendor/github.com/golang/mock/mockgen
//go:generate go build -v -o=${ROOT}bin/protoc-gen-grpc-gateway ./vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
//go:generate go build -v -o=${ROOT}bin/protoc-gen-swagger ./vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
//go:generate go build -v -o=${ROOT}bin/gex ./vendor/github.com/izumin5210/gex/cmd/gex
//go:generate go build -v -o=${ROOT}bin/golint ./vendor/golang.org/x/lint/golint

