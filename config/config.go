package config

import (
	"github.com/joho/godotenv"
)

type (
	Config struct {
		*App
		*HTTP
		*DB
	}

	App struct {
		Name string
	}

	HTTP struct {
		Port string
	}

	DB struct {
		URL string
	}
)

func NewConfigFrom() (*Config, error) {
	// path, err := os.Getwd()

	// if err != nil {
	// 	return nil, err
	// }

	// fmt.Println("PATH TO ENV", path+"/config/.env")

	envs, err := godotenv.Read("./.env")
	if err != nil {
		return nil, err
	}

	return &Config{
		&App{
			Name: envs["APP_NAME"],
		},
		&HTTP{
			Port: envs["SERVER_PORT"],
		},
		&DB{
			URL: envs["DATABASE_URL"],
		},
	}, nil
}
