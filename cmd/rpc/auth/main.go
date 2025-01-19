package auth

import (
	"project/common/cmd"
	"project/common/config"
	rpcserver "project/internal/auth/delivery/rpc-server"
)

func main() {
	authCmd := cmd.NewRpcCmd("auth")
	authCmd.PreLoadConfig()
	authCmd.SetSvcName(config.Config.RPCRegisterName.MicroCleanAuthName)

	authCmd.AddPortFlag()
	authCmd.AddPrometheusPortFlag()
	if err := authCmd.Exec(); err != nil {
		panic(err.Error())
	}
	if err := authCmd.StartSvr(config.Config.RPCRegisterName.MicroCleanAuthName, rpcserver.StartAuthServer); err != nil {
		panic(err.Error())
	}
}
