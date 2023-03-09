package configs

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type LineService struct {
	ChannelToken  string `yaml:"channel_token"`
	ChannelSecret string `yaml:"channel_secret"`
}

type ChatGPTService struct {
	Key string `yaml:"key"`
}

type Services struct {
	LineService    `yaml:"line_service"`
	ChatGPTService `yaml:"chatgpt_service"`
}

type Config struct {
	Server   Server   `yaml:"server"`
	Services Services `yaml:"services"`
}

func ReadEnv(fileName string) (*Config, error) {
	yamlFile, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	var config Config
	if err = yaml.Unmarshal(yamlFile, &config); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &config, nil
}
