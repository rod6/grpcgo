package service

import (
	"context"

	"github.com/libgo/logx"

	"demo/pkg/ping"
)

type PingService struct {
	ping.UnimplementedPingServiceServer
}

func (*PingService) Pong(ctx context.Context, in *ping.Ping) (*ping.Ping, error) {
	logx.Infof("rec: %s", in)
	return in, nil
}
