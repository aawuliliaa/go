package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	pb "project/api"
	"project/configs"
	"project/internal/log"
	"project/internal/service"
)

func NewGrpcServer() (*grpc.Server, error) {
	grpcPort := configs.ProjectConfig.GrpcPort
	grpcIp := configs.ProjectConfig.GrpcIp
	//地址
	grpcAddr := fmt.Sprintf("%s:%s", grpcIp, grpcPort)
	//1.监听
	listener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		return nil, fmt.Errorf("grpc listen err,err is%v", err)
	}
	fmt.Printf("grpc监听端口 %s\n", grpcAddr)
	//2.实例化gRPC
	grpcServer := grpc.NewServer()
	//3.在gRPC上注册微服务
	//
	pb.RegisterGroupServiceServer(grpcServer, &service.GroupService{})

	//4.启动服务端
	go func() {
		err = grpcServer.Serve(listener)
		log.Logger.Infof("start grpc server err,err is%v", err)

	}()

	return grpcServer, nil
}
