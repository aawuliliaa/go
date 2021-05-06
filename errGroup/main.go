package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	errG, errGroupCtx := errgroup.WithContext(context.Background())
	r := gin.Default()
	serverStop := make(chan string)
	v1Group := r.Group("/api/v1")
	v1Group.GET("/ping", func(ginCtx *gin.Context) {
		ginCtx.JSON(200, gin.H{
			"message": "success",
		})
	})
	v1Group.GET("/stop", func(ginCtx *gin.Context) {
		serverStop <- "stop"
	})
	server := http.Server{Addr: ":8080", Handler: r}
	//g1
	//g1 退出了，所有协程都将退出，因为
	//g1退出后，errGroupCtx.Done()将不再阻塞，g2,g3也会 随之退出
	//mian函数中的g.wait()退出，所有协程都会退出
	errG.Go(func() error {
		return server.ListenAndServe()
	})
	//g2
	//g2 退出了，所有协程都将退出，因为
	//g2退出时，调用了shutdown,g1会退出
	//g2退出时，errGroupCtx.Done()将不再阻塞，g3也会 随之退出
	//mian函数中的g.wait()退出，所有协程都会退出
	errG.Go(func() error {
		select {
		case <-errGroupCtx.Done():
			log.Println("g2 errGroupCtx.Done,errGroup exit")
		case <-serverStop:
			log.Println("g2 server stop,errGroup exit")
		}
		timeoutCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		// 这里不是必须的，但是如果使用 _ 的话静态扫描工具会报错，加上也无伤大雅
		defer cancel()
		log.Println("g2 server shutdown")
		return server.Shutdown(timeoutCtx)
	})
	//g3
	//g3 退出了，所有协程都将退出，因为
	//g3退出时，errGroupCtx.Done()将不再阻塞，g2也会 随之退出
	//g2退出时，调用了shutdown,g1会退出
	//mian函数中的g.wait()退出，所有协程都会退出
	errG.Go(func() error {
		quit := make(chan os.Signal, 0)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-errGroupCtx.Done():
			log.Printf("g3 errGroupCtx.Done, errGroup exit, errGroupCtx.Err():%v", errGroupCtx.Err())
			return errGroupCtx.Err()
		case sig := <-quit:
			log.Printf("g3 get os signal: %v", sig)
			return fmt.Errorf("g3 get os signal: %v", sig)

		}
	})
	fmt.Printf("main errgroup exiting: %+v\n", errG.Wait())

}
