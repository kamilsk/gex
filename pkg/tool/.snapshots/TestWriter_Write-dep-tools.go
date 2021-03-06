// Code generated by github.com/kamilsk/egg. DO NOT EDIT.

// +build tools

package tools

import (
	_ "github.com/gogo/protobuf/protoc-gen-gogofast"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger"
	_ "github.com/volatiletech/sqlboiler"
	_ "github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql"
)

//go:generate go build -v -o=${ROOT}bin/protoc-gen-gogofast ./vendor/github.com/gogo/protobuf/protoc-gen-gogofast
//go:generate go build -v -o=${ROOT}bin/protoc-gen-grpc-gateway ./vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
//go:generate go build -v -o=${ROOT}bin/protoc-gen-swagger ./vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
//go:generate go build -v -o=${ROOT}bin/sqlboiler ./vendor/github.com/volatiletech/sqlboiler
//go:generate go build -v -o=${ROOT}bin/sqlboiler-psql ./vendor/github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql

