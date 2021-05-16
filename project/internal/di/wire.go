// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"github.com/google/wire"
	"project/internal/server/grpc"
	"project/internal/server/http"
)
//go:generate project t wire
func InitApp() ( *App,func(), error) {
	panic(wire.Build( http.NewHttpServer, grpc.NewGrpcServer,NewApp))
}