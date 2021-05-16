package di

import (
	"context"
	"google.golang.org/grpc"
	"net/http"
	"project/internal/log"
	"time"
)

type App struct {
	HttpServer *http.Server
	GrpcServer *grpc.Server
}

func NewApp(httpServer *http.Server, grpcServer *grpc.Server) (*App,func(),error) {
	app := &App{
		HttpServer: httpServer,
		GrpcServer: grpcServer,
	}
	closeFunc := func() {
		ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
		if err := httpServer.Shutdown(ctx); err != nil {
			log.Logger.Errorf("grpcSrv.Shutdown error(%v)", err)
		}
		grpcServer.GracefulStop();
		cancel()
	}
	return app,closeFunc,nil
}
