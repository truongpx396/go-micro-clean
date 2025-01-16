package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const (
	FileName             = "config.yaml"
	NotificationFileName = "notification.yaml"
	DefaultFolderPath    = "../config/"
)

type configStruct struct {
	Zookeeper struct {
		Schema   string   `yaml:"schema"`
		Address  []string `yaml:"address"`
		Username string   `yaml:"username"`
		Password string   `yaml:"password"`
	} `yaml:"zookeeper"`
	Postgre struct {
		Address       []string `yaml:"address"`
		Username      string   `yaml:"username"`
		Password      string   `yaml:"password"`
		Database      string   `yaml:"database"`
		MaxOpenConn   int      `yaml:"maxOpenConn"`
		MaxIdleConn   int      `yaml:"maxIdleConn"`
		MaxLifeTime   int      `yaml:"maxLifeTime"`
		LogLevel      int      `yaml:"logLevel"`
		SlowThreshold int      `yaml:"slowThreshold"`
	} `yaml:"postgre"`
	Redis struct {
		Address  []string `yaml:"address"`
		Username string   `yaml:"username"`
		Password string   `yaml:"password"`
	} `yaml:"redis"`
	RPC struct {
		RegisterIP string `yaml:"registerIP"`
		ListenIP   string `yaml:"listenIP"`
	} `yaml:"rpc"`
	API struct {
		MicroCleanApiPort []int  `yaml:"microCleanApiPort"`
		ListenIP          string `yaml:"listenIP"`
	} `yaml:"api"`
	RPCPort struct {
		MicroCleanUserPort []int `yaml:"microCleanUserPort"`
		MicroCleanItemPort []int `yaml:"microCleanItemPort"`
		MicroCleanAuthPort []int `yaml:"microCleanAuthPort"`
		MicroCleanPushPort []int `yaml:"microCleanPushPort"`
	} `yaml:"rpcPort"`
	RPCRegisterName struct {
		MicroCleanUserName string `yaml:"microCleanUserName"`
		MicroCleanItemName string `yaml:"microCleanItemName"`
		MicroCleanAuthName string `yaml:"microCleanAuthName"`
		MicroCleanPushName string `yaml:"microCleanPushName"`
	} `yaml:"rpcRegisterName"`
	Log struct {
		StorageLocation     string `yaml:"storageLocation"`
		RotationTime        int    `yaml:"rotationTime"`
		RemainRotationCount int    `yaml:"remainRotationCount"`
		RemainLogLevel      int    `yaml:"remainLogLevel"`
		IsStdout            bool   `yaml:"isStdout"`
		IsJson              bool   `yaml:"isJson"`
		WithStack           bool   `yaml:"withStack"`
	} `yaml:"log"`
	Push struct {
		Enable string `yaml:"enable"`
		FCM    struct {
			ServiceAccount string `yaml:"serviceAccount"`
		} `yaml:"fcm"`
		JPNS struct {
			AppKey       string `yaml:"appKey"`
			MasterSecret string `yaml:"masterSecret"`
			PushUrl      string `yaml:"pushUrl"`
			PushIntent   string `yaml:"pushIntent"`
		} `yaml:"jpns"`
	} `yaml:"push"`
	Secret      string `yaml:"secret"`
	TokenPolicy struct {
		Expire int `yaml:"expire"`
	} `yaml:"tokenPolicy"`
	Prometheus struct {
		Enable             bool  `yaml:"enable"`
		UserPrometheusPort []int `yaml:"userPrometheusPort"`
		ItemPrometheusPort []int `yaml:"itemPrometheusPort"`
		AuthPrometheusPort []int `yaml:"authPrometheusPort"`
		PushPrometheusPort []int `yaml:"pushPrometheusPort"`
	} `yaml:"prometheus"`
}

var Config *configStruct

func LoadConfig(configName, configFolderPath string) error {

	if configFolderPath == "" {
		configFolderPath = DefaultFolderPath
	}
	configPath := filepath.Join(configFolderPath, configName)
	defer func() {
		fmt.Println("use config", configPath)
	}()

	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		log.Fatalf("Config file does not exist: %v", err)
	}

	file, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// var cfg Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(Config); err != nil {
		return err
	}

	return nil
}
