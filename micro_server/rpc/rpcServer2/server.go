package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 结构体，用于注册的
type Arith struct{}

// 声明参数结构体
type ArithRequest struct {
	A, B int
}

// 返回给客户端的结果
type ArithResponse struct {
	// 乘积
	Pro int
	// 商
	Quo int
	// 余数
	Rem int
}

// 乘法
func (this *Arith) Multiply(req ArithRequest, res *ArithResponse) error {
	res.Pro = req.A * req.B
	return nil
}

// 商和余数
func (this *Arith) Divide(req ArithRequest, res *ArithResponse) error {
	if req.B == 0 {
		return errors.New("除数不能为0")
	}
	// 除
	res.Quo = req.A / req.B
	// 取模
	res.Rem = req.A % req.B
	return nil
}

// 主函数
func main() {
	rpc.Register(new(Arith))
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	// 循环监听服务
	for {
		conn, err := lis.Accept()
		if err != nil {
			continue
		}
		go func(conn net.Conn) {
			fmt.Println("new client")
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}
