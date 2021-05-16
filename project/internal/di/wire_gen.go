// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package di

import (
	"project/internal/server/grpc"
	"project/internal/server/http"
)

// Injectors from wire.go:

//go:generate project t wire
func InitApp() (*App, func(), error) {
	server, err := http.NewHttpServer()
	if err != nil {
		return nil, nil, err
	}
	grpcServer, err := grpc.NewGrpcServer()
	if err != nil {
		return nil, nil, err
	}
	app, cleanup, err := NewApp(server, grpcServer)
	if err != nil {
		return nil, nil, err
	}
	return app, func() {
		cleanup()
	}, nil
}
