package cmd

import (
	"context"
	"errors"
	"project/common/startrpc"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

type RpcCmd struct {
	*RootCmd
}

func NewRpcCmd(name string) *RpcCmd {
	authCmd := &RpcCmd{NewRootCmd(name)}
	return authCmd
}

func (a *RpcCmd) Exec() error {
	a.Command.Run = func(cmd *cobra.Command, args []string) {
		a.port = a.getPortFlag(cmd)
		a.prometheusPort = a.getPrometheusPortFlag(cmd)
	}
	return a.Execute()
}

func (a *RpcCmd) StartSvr(
	name string,
	rpcFn func(ctx context.Context, server *grpc.Server) error,
) error {
	if a.GetPortFlag() == 0 {
		return errors.New("port is required")
	}
	return startrpc.Start(a.GetPortFlag(), name, a.GetPrometheusPortFlag(), rpcFn)
}
