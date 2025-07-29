package config

import (
	"fmt"
	"os"

	"github.com/youngprinnce/geolocation-service/internal/logger"
	"gopkg.in/yaml.v2"
)

type App struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Env     string `yaml:"env"`
}

type Database struct {
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	DbName   string `yaml:"db_name"`
}

type Server struct {
	Listen string `yaml:"listen"`
}

type Config struct {
	App      App      `yaml:"app"`
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
}

var conf Config

func LoadConfig(path string) *Config {
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		logger.Fatal(fmt.Sprintf("yamlFile.Get err   #%v ", err))
	}

	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Unmarshal: %v", err))
	}

	return &conf
}

func GetConfig() *Config {
	return &conf
}
