package config

import (
	"fmt"
	"os"
)

const (
	envHTTP   = "EKOLO_HTTP"
	envDBHost = "EKOLO_DB_HOST"
	envDBPort = "EKOLO_DB_PORT"
	envDBName = "EKOLO_DB_NAME"
	envDBUser = "EKOLO_DB_USER"
	envDBPass = "EKOLO_DB_PASS"
)

type Config struct {
	HTTPAddr string
	DBHost   string
	DBPort   string
	DBUser   string
	DBPass   string
	DBName   string
}

func (cfg Config) GetDBDSN() string {
	dsn := "host=%s port=%s dbname=%s user='%s' password=%s sslmode=disable"
	return fmt.Sprintf(dsn, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBUser, cfg.DBPass)
}

func getValue(envKey string) string {
	var val string
	if v := os.Getenv(envKey); v != "" {
		val = v
	}
	return val
}

func New() Config {
	var cfg = Config{
		HTTPAddr: ":8080",
		DBPort:   "5432",
	}
	if v := getValue(envHTTP); v != "" {
		cfg.HTTPAddr = v
	}
	cfg.DBHost = getValue(envDBHost)
	if p := getValue(envDBPort); p != "" {
		cfg.DBPort = p
	}
	cfg.DBUser = getValue(envDBUser)
	cfg.DBPass = getValue(envDBPass)
	cfg.DBName = getValue(envDBName)
	return cfg
}
