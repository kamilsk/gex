// Code generated by github.com/kamilsk/egg. DO NOT EDIT.

// +build tools

package tools

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"
)

//go:generate go build -v -o=${ROOT}bin/protoc-gen-grpc-gateway ./vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
