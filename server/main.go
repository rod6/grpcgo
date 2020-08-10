package main

import (
	"github.com/libgo/micro"
	"google.golang.org/grpc"

	"demo/pkg/ping"
	"demo/pkg/user"
	"demo/server/service"
)

func main() {
	m := micro.New("hello-world")

	s := grpc.NewServer()
	ping.RegisterPingServiceServer(s, &service.PingService{})
	user.RegisterUserServiceServer(s, &service.UserService{})

	m.ServeGRPC(":9090", s)

	m.Start()
}
