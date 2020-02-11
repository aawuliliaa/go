package main

import (
	"context"
	"fmt"
	"github.com/go-micro"
	"github.com/go-micro/errors"
	pb "go_5/my-micro/02/proto"
	"log"
)

type Example struct {
}

type Foo struct {
}

func (e *Example) Call(ctx context.Context, req *pb.CallRquest, rsp *pb.CallResponse) error {
	log.Print("收到Example.Call请求")
	if len(req.Name) == 0 {
		return errors.BadRequest("go.micro.api.example", "no name")
	}
	rsp.Message = "Exmaple.Call接收到了你的请求" + req.Name
	return nil
}

func (f *Foo) Bar(ctx context.Context, req *pb.EmptyRquest, rsp *pb.EmptyResponse) error {
	log.Print("收到Foo.Bar请求")
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.example"),
	)
	service.Init()
	err := pb.RegisterExampleHandler(service.Server(), new(Example))
	if err != nil {
		fmt.Println(err)
	}
	err = pb.RegisterFooHandler(service.Server(), new(Foo))
	if err != nil {
		fmt.Println(err)
	}
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
