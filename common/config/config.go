package config

var Config *configStruct

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
		RotationSize        int    `yaml:"rotationSize"`
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
