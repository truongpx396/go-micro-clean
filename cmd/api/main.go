package main

import (
	"context"
	"fmt"
	"net"
	"go-micro-clean/common/cmd"
	"go-micro-clean/common/config"
	"go-micro-clean/internal/api"
	"strconv"

	_ "net/http/pprof"

	"go-micro-clean/tools/log"
)

//swag init --parseInternal --pd --dir cmd/openim-api,internal/api/ -g ../../internal/api/route.go --output cmd/openim-api/docs

func main() {
	apiCmd := cmd.NewApiCmd()
	apiCmd.AddApi(run)
	if err := apiCmd.Execute(); err != nil {
		panic(err.Error())
	}
}

func run(port int) error {
	if port == 0 {
		return fmt.Errorf("port is empty")
	}
	// rdb, err := cache.NewRedis()
	// if err != nil {
	// 	return err
	// }
	// fmt.Println("api start init discov client")
	// var client discoveryregistry.SvcDiscoveryRegistry
	// client, err = openkeeper.NewClient(config.Config.Zookeeper.ZkAddr, config.Config.Zookeeper.Schema,
	// 	openkeeper.WithFreq(time.Hour), openkeeper.WithUserNameAndPassword(
	// 		config.Config.Zookeeper.Username,
	// 		config.Config.Zookeeper.Password,
	// 	), openkeeper.WithRoundRobin(), openkeeper.WithTimeout(10), openkeeper.WithLogger(log.NewZkLogger()))
	// if err != nil {
	// 	return err
	// }

	// fmt.Println("api start CreateRpcRootNodes: ", config.Config.GetServiceNames())
	// if err := client.CreateRpcRootNodes(config.Config.GetServiceNames()); err != nil {
	// 	return err
	// }
	// fmt.Println("api register public config to discov")
	// if err := client.RegisterConf2Registry(constant.OpenIMCommonConfigKey, config.Config.EncodeConfig()); err != nil {
	// 	return err
	// }

	ctx := context.Background()

	fmt.Println("api register public config to discov success")
	router := api.NewGinRouter(ctx)
	fmt.Println("api init router success")

	var address string
	if config.Config.API.ListenIP != "" {
		address = net.JoinHostPort(config.Config.API.ListenIP, strconv.Itoa(port))
	} else {
		address = net.JoinHostPort("0.0.0.0", strconv.Itoa(port))
	}

	// fmt.Println("start api server, address: ", address, ", OpenIM version: ", config.Config.Version)
	log.Info("start server success", "address", address, "version", config.Version)
	err := router.Run(address)
	if err != nil {
		log.Error("api run failed ", err, "address", address)
		return err
	}
	return nil
}
