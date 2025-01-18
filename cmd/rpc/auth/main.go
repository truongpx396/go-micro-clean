package auth

import "project/common/config"

func main() {
	authCmd := cmd.NewRpcCmd("auth")
	authCmd.PreLoadConfig()
	authCmd.SetSvcName(config.Config.RPCRegisterName.MicroCleanAuthName)

	authCmd.AddPortFlag()
	authCmd.AddPrometheusPortFlag()
	if err := authCmd.Exec(); err != nil {
		panic(err.Error())
	}
	if err := authCmd.StartSvr(config.Config.RPCRegisterName.MicroCleanAuthName, auth.StartAuthServer); err != nil {
		panic(err.Error())
	}
}
