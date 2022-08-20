package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("config/config.env"); err != nil {
		log.Panic(err)
	}
}

type ApplicationConfig struct {
	DB       database
	HttpPort string `env:"HTTP_PORT"`
	GrpcPort string `env:"GRPC_PORT"`
}

func GetConfig() (ApplicationConfig, error) {
	conf := ApplicationConfig{}
	return conf, env.Parse(&conf)
}

type database struct {
	Name     string `env:"PG_NAME"`
	Host     string `env:"PG_HOST"`
	User     string `env:"PG_USER"`
	Password string `env:"PG_PASSWORD"`
	Port     string `env:"PG_PORT"`
}

func (db database) GetDSN() string {
	return fmt.Sprintf("postgresql://%v:%v@%v:%v/postgres?search_path=%v&sslmode=disable",
		db.User,
		db.Password,
		db.Host,
		db.Port,
		db.Name)
}
