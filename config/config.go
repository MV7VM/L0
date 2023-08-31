package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

// urlExample := "postgres://username:password@localhost:5432/database_name"
// cleanenv
type Storage struct {
	Username string `json:"username,omitempty" env:"DB_Username"`
	Password string `json:"password,omitempty" env:"DB_Password"`
	Host     string `json:"host,omitempty" env:"DB_Host"`
	Port     string `json:"port,omitempty" env:"DB_Port"`
	Name     string `json:"db___name,omitempty" env:"DB_Name"`
}

type Config struct {
	DB Storage
}

// Config func to get env value
func Config_load() Config {
	var cfg Config
	err := cleanenv.ReadConfig(".env", &cfg)
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return cfg
}
