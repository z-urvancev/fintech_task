package config

import (
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

type Config struct {
	Port string   `yaml:"port"`
	DB   DBConfig `yaml:"db"`
}

type DBConfig struct {
	Username string `yaml:"username"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"dbname"`
	Password string `yaml:"password"`
	SSLMode  string `yaml:"sslmode"`
}

func InitConfig(config *Config) error {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		return envErr
	}

	filename, fileErr := filepath.Abs("./config/config.yml")
	if fileErr != nil {
		return fileErr
	}

	yamlFile, yamlErr := os.ReadFile(filename)
	if yamlErr != nil {
		return yamlErr
	}

	marshalErr := yaml.Unmarshal(yamlFile, config)
	if marshalErr != nil {
		return marshalErr
	}

	config.DB.Password = os.Getenv("DB_PASSWORD")
	return nil
}
