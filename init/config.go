package init

import "os"

const configDirectory = "./../config/"

type Configuration struct {
	host string
	port int8
}

func (c *Configuration) Host() string {
	return c.host
}

func (c *Configuration) Port() int8 {
	return c.port
}

func Create() {
 	configFile, err := os.OpenFile(configurationFilePath(), 0, 644)
	if err != nil {
		println(err.Error())
	}
	defer configFile.Close()
}

func isProductionEnvironment() bool {
	return os.Getenv("debug") != "true"
}

func configurationFilePath() string {
	configFileName := "config.dev.json"
	if isProductionEnvironment() == true {
		configFileName = "config.prod.json"
	}
	return configDirectory + configFileName
}