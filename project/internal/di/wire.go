// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"github.com/google/wire"
	"project/internal/dao"
	"project/internal/server/grpc"
	"project/internal/server/http"
	"project/internal/service"
)
//go:generate
func InitApp() ( *App,func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.NewHttpServer, grpc.NewGrpcServer,NewApp))
}