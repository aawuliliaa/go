package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	pb "go_5/my-micro/01/proto"
	"log"
)

// 声明结构体
type Hello struct{}

func (g *Hello) Info(ctx context.Context, req *pb.InfoRequest, rep *pb.InfoResponse) error {
	rep.Msg = "你好" + req.Username
	return nil
}

func main() {
	// 1.得到服务端实例
	service := micro.NewService(
		// 设置微服务的名，用来访问
		// micro call hello Hello.Info {"username":"zhangsan"}
		micro.Name("hello"),
	)
	// 2.初始化
	service.Init()
	// 3.服务注册
	err := pb.RegisterHelloHandler(service.Server(), new(Hello))
	if err != nil {
		fmt.Println(err)
	}
	// 4.启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
