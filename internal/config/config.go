package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type MySQL struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type Config struct {
	MySQL MySQL `yaml:"mysql"`
	Redis struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Password string `yaml:"password"`
	} `yaml:"redis"`
	OpenAI struct {
		APIKey string `yaml:"api_key"`
	} `yaml:"openai"`
}

var AppConfig Config

func LoadConfig(configPath string) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("无法读取配置文件: %v", err)
	}

	err = yaml.Unmarshal(data, &AppConfig)
	if err != nil {
		log.Fatalf("无法解析配置文件: %v", err)
	}
}
