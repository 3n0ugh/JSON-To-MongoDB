package config

import (
	"os"

	"github.com/golobby/dotenv"
)

type Config struct {
	DatabaseName    string `env:"DB_NAME"`
	DatabasePort    string `env:"DB_PORT"`
	DatabaseHost    string `env:"DB_HOST"`
	DatabaseColName string `env:"DB_COL_NAME"`

	JsonPath string `env:"JSON_PATH"`
}

func GetConfig() (*Config, error) {
	// initialize the config struct
	config := &Config{}

	// open dotenv file
	file, err := os.Open(".env")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// adding data to config struct
	err = dotenv.NewDecoder(file).Decode(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
