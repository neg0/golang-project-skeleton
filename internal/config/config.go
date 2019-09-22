package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

const directoryPath = "./../../config"

type API struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type Configuration struct {
	API API `json:"api"`
}

var instance *Configuration

func Instance() (*Configuration, error) {
	var hasError error
	once := sync.Once{}
	once.Do(func() {
		cnf, err := newConfig().fromFile(nil)
		if err != nil {
			hasError = err
			return
		}
		instance = cnf
	})
	return instance, hasError
}

// Creates Config from JSON config file by un-marshalling the contents
func (c *Configuration) fromFile(path []byte) (*Configuration, error) {
	if path == nil {
		path = filePath()
	}

	configFile, err := ioutil.ReadFile(string(path))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(configFile, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func newConfig() *Configuration {
	return &Configuration{}
}

// Determines if environment is production of development
func isProd() bool {
	return os.Getenv("debug") != "true"
}

// It returns the configuration file path
func filePath() []byte {
	configFileName := "config.dev.json"
	if isProd() {
		configFileName = "config.prod.json"
	}
	return []byte(fmt.Sprintf("%s/%s", directoryPath, configFileName))
}
