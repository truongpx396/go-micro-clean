package cmd

import (
	"context"
	"errors"
	"go-micro-clean/common/config"
	"go-micro-clean/common/startrpc"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

type rpcCmd struct {
	*RootCmd
}

func NewRpcCmd(name string) *rpcCmd {
	authCmd := &rpcCmd{NewRootCmd(name)}
	return authCmd
}

func (a *rpcCmd) Exec() error {
	a.Command.Run = func(cmd *cobra.Command, args []string) {
		a.initRpcServiceNameByConfig()
		a.port = a.getPortFlag(cmd)
		a.prometheusPort = a.getPrometheusPortFlag(cmd)
	}
	return a.Execute()
}

func (r *rpcCmd) initRpcServiceNameByConfig() {
	switch r.Name {
	case config.AuthName:
		r.Name = config.Config.RPCRegisterName.MicroCleanAuthName
	case config.UserName:
		r.Name = config.Config.RPCRegisterName.MicroCleanUserName
	case config.ItemName:
		r.Name = config.Config.RPCRegisterName.MicroCleanItemName
	case config.PushName:
		r.Name = config.Config.RPCRegisterName.MicroCleanPushName
	default:
		r.Name = "rpc-server"
	}
}

func (a *rpcCmd) StartSvr(
	rpcFn func(ctx context.Context, server *grpc.Server) error,
) error {
	if a.getPort() == 0 {
		return errors.New("port is required")
	}
	return startrpc.Start(a.Name, a.getPort(), a.getPrometheusPort(), rpcFn)
}
