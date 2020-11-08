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
	WorkHost                string `yaml:"workHost"`
	WorkPort                int    `yaml:"workPort"`
	*MySQLConfig            `yaml:"mysql"`
	*NSQConfig              `yaml:"nsqProducer"`
	*ResetPwdConsumerConfig `yaml:"nsqConsumer"`
	*MailConfig             `yaml:"mail"`
	*RedisConfig            `yaml:"redis"`
}

type MySQLConfig struct {
	User     string `yaml:"user"`
	PassWord string `yaml:"pwd"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DB       string `yaml:"db"`
}

type NSQConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type ResetPwdConsumerConfig struct {
	Topic   string `yaml:"topic"`
	Channel string `yaml:"channel"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
}

type MailConfig struct {
	Account string `yaml:"account"`
	Pwd     string `yaml:"pwd"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DB       int    `yaml:"db"`
	PoolSize int    `yaml:"poolSize"`
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
