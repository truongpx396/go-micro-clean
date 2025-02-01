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
	DefaultFolderPath    = "config/"
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

	rootDir, err := findProjectRoot()
	if err != nil {
		log.Fatalf("Error finding project root: %v", err)
	}

	configPath := filepath.Join(rootDir, configFolderPath, configName)

	defer func() {
		fmt.Println("use config", configPath)
	}()

	_, err = os.Stat(configPath)
	if os.IsNotExist(err) {
		log.Fatalf("Config file does not exist: %v", err)
	}

	file, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	Config = &configStruct{}

	// var cfg Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(Config); err != nil {
		return err
	}

	return nil
}

func findProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		// Check for go.mod or .git directory to detect the project root
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		if _, err := os.Stat(filepath.Join(dir, ".git")); err == nil {
			return dir, nil
		}

		// Move up one directory
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("project root not found")
		}
		dir = parent
	}
}
