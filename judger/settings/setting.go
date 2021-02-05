package settings

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

var Conf = new(Config)

type Config struct {
	Port         string `yaml:"port"`
	MaxTask      int    `yaml:"maxTask"`
	*MySQLConfig `yaml:"mysql"`
}

type MySQLConfig struct {
	User     string `yaml:"user"`
	PassWord string `yaml:"pwd"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DB       string `yaml:"db"`
}

func Init() {
	basePath, err := os.Getwd()
	if err != nil {
		fmt.Println("base path error.")
	}
	fileName := filepath.Join(basePath, "conf", "config.yaml")
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("load config file error.")
	}
	err = yaml.Unmarshal(yamlFile, Conf)
	if err != nil {
		fmt.Printf("load mysql config fail: %v\n", err.Error())
	}
}
