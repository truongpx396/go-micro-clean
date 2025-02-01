package main

import (
	"go-micro-clean/common/cmd"
	"go-micro-clean/common/config"
	rpcserver "go-micro-clean/internal/auth/delivery/rpc-server"
)

func main() {
	authCmd := cmd.NewRpcCmd(config.AuthName)

	if err := authCmd.Exec(); err != nil {
		panic(err.Error())
	}
	if err := authCmd.StartSvr(rpcserver.StartAuthServer); err != nil {
		panic(err.Error())
	}
}
