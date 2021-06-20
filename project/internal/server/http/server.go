package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	pb "project/api"
	"project/internal/log"
	"project/internal/service"
)
var AppService *service.Service
func NewHttpServer(appService *service.Service) (*http.Server,error) {
	engine := gin.Default()
	engine = InitRouters(engine)
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Logger.Fatalf("httpServer listen: %s\n", err)
		}
	}()
	AppService = appService

	return httpServer,nil
}
func InitRouters(engine *gin.Engine) *gin.Engine {
	base := engine.Group("/api/v1")
	{
		base.GET("/getGroup", func(ctx *gin.Context) {
			res,err :=AppService.GroupService.GetGroupInfo(context.Background(),&pb.GetGroupRequest{Name: "vita"})
			if err!=nil{
				ctx.JSON(500, gin.H{
					"message": "get err",
				})
				return
			}
			ctx.JSON(200, gin.H{
				"name": res.Name,
			})
		})

	}
	return engine
}
