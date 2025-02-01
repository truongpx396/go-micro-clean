package main

import (
	"go-micro-clean/common/cmd"
	"go-micro-clean/common/config"
	rpcserver "go-micro-clean/internal/user/delivery/rpc-server"
)

func main() {
	userCmd := cmd.NewRpcCmd(config.UserName)

	if err := userCmd.Exec(); err != nil {
		panic(err.Error())
	}
	if err := userCmd.StartSvr(rpcserver.StartUserServer); err != nil {
		panic(err.Error())
	}
}
