package main

import (
	"github.com/gin-gonic/gin"
	//"golang.org/x/sync/errgroup"
	"net/http"
)

func ping(ginCtx *gin.Context)  {
	ginCtx.JSON(200, gin.H{
		"message": "success",
	})
}


func main()  {
	//errGroup,errGroupCtx := errgroup.WithContext(context.Background())
	r :=gin.Default()
	v1Group:=r.Group("/api/v1")
	v1Group.GET("/ping",ping)
	server := http.Server{Addr:":8080",Handler:r}
	_ = server.ListenAndServe()
}
