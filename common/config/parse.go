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

var Version string

const (
	UserName = "user"
	ItemName = "item"
	AuthName = "auth"
	PushName = "push"
)

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
