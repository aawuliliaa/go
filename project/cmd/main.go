package main

import (
	"os"
	"os/signal"
	"project/configs"
	"project/internal/di"
	"project/internal/log"
	"syscall"
	"time"
)

func main()  {
	log.InitLogger()
	err := configs.LoadConfig("")
	if err != nil {
		panic(err)
	}
	_, closeFunc, err := di.InitApp()
	if err != nil {
		panic(err)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Logger.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeFunc()
			log.Logger.Info("demo exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
