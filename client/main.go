package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/libgo/logx"
	"google.golang.org/grpc"

	"demo/pkg/ping"
)

var pingCli = func() ping.PingServiceClient {
	cc, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return ping.NewPingServiceClient(cc)
}()

func main() {
	for {
		resp, err := pingCli.Pong(context.Background(), &ping.Ping{Id: uuid.New().String()})
		if err != nil {
			logx.Errorf("error: %s", err.Error())
		}
		fmt.Println(resp)
		<-time.After(time.Second)
	}
}
