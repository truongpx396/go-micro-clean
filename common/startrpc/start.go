package startrpc

import (
	"context"
	"fmt"
	"net"
	"go-micro-clean/common/config"
	"strconv"

	"google.golang.org/grpc"
)

func Start(
	rpcRegisterName string,
	rpcPort int,
	prometheusPort int,
	rpcFn func(ctx context.Context, server *grpc.Server) error,
	options ...grpc.ServerOption,
) error {
	fmt.Println(
		"start",
		rpcRegisterName,
		"server, port: ",
		rpcPort,
		"prometheusPort:",
		prometheusPort,
		", MicroClean version: ",
		config.Version,
	)
	listener, err := net.Listen(
		"tcp",
		net.JoinHostPort("", strconv.Itoa(rpcPort)),
	)
	if err != nil {
		return err
	}
	defer listener.Close()

	// registerIP, err := network.GetRpcRegisterIP(config.Config.RPC.RegisterIP)
	// if err != nil {
	// 	return err
	// }
	// ctx 中间件
	if config.Config.Prometheus.Enable {
		// prome.NewGrpcRequestCounter()
		// prome.NewGrpcRequestFailedCounter()
		// prome.NewGrpcRequestSuccessCounter()
		// unaryInterceptor := mw.InterceptChain(grpcprometheus.UnaryServerInterceptor, mw.RpcServerInterceptor)
		// options = append(options, []grpc.ServerOption{
		// 	grpc.StreamInterceptor(grpcprometheus.StreamServerInterceptor),
		// 	grpc.UnaryInterceptor(unaryInterceptor),
		// }...)
	} else {
		// options = append(options, grpc.ChainUnaryInterceptor(RpcServerInterceptor))
	}
	srv := grpc.NewServer(options...)
	defer srv.GracefulStop()
	err = rpcFn(context.Background(), srv)
	if err != nil {
		return err
	}

	go func() {
		if config.Config.Prometheus.Enable && prometheusPort != 0 {
			// prome.Enable = true
			// if err := prome.StartPrometheusSrv(prometheusPort); err != nil {
			// 	panic(err.Error())
			// }
		}
	}()
	return srv.Serve(listener)
}
