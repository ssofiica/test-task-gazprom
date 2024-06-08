package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type (
	Project struct {
		Server   `yaml:"server"`
		Postgres `yaml:"postgresql"`
		Redis    `yaml:"redis"`
	}

	Server struct {
		Host             string `yaml:"host"    env:"SRV_HOST"`
		Port             uint16 `yaml:"port"    env:"SRV_PORT"`
		WriteTimeout     uint16 `yaml:"write-timeout"    env:"SRV_WRITE_TM"`
		ReadTimeout      uint16 `yaml:"read-timeout"    env:"SRV_READ_TM"`
		IdleTimeout      uint16 `yaml:"idle-timeout"    env:"SRV_IDLE_TM"`
		ShutdownDuration uint16 `yaml:"shutdown-duration"    env:"SRV_SHUTDOWN_DUR"`
	}

	Postgres struct {
		Host            string `yaml:"host"    env:"PG_HOST"`
		Port            uint16 `yaml:"port" env:"PG_PORT"`
		User            string `yaml:"user" env:"PG_USER"`
		Password        string `yaml:"password" env:"PG_PASSWORD"`
		Database        string `yaml:"database" env:"PG_DB"`
		SslMode         string `yaml:"sslmode" env:"PG_SSL_MODE"`
		MaxOpenConns    uint32 `yaml:"max-open-connections" env:"PG_MAX_OPEN_CONN"`
		ConnMaxLifetime uint16 `yaml:"conn-max-lifetime" env:"PG_CONN_MAX_LIFETIME"`
		MaxIdleConns    uint32 `yaml:"max-idle-conns" env:"PG_IDLE_CONNS"`
		ConnMaxIdleTime uint16 `yaml:"conn-max-idle-time" env:"PG_MAX_IDLE_TIME"`
	}

	Redis struct {
		Host            string `yaml:"host"    env:"R_HOST"`
		Port            uint16 `yaml:"port" env:"R_PORT"`
		DatabaseSession int    `yaml:"database-session" env:"R_DB_SESSION"`
		User            string `yaml:"user" env:"R_USER"`
		Password        string `yaml:"password" env:"R_PASSWORD"`
	}
)

func NewConfig() *Project {
	cfg := &Project{}

	buf, err := os.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatal("Error reading application configuration: ", err.Error())
	}

	err = yaml.Unmarshal(buf, cfg)
	if err != nil {
		log.Fatal("Error creating configuration object: ", err.Error())
	}

	fmt.Println("Reading configuration successful")
	return cfg
}
