package config

import (
	"os"
	"strconv"
)

type PGConfig struct {
	DBUser        string
	DBPassword    string
	DBHost        string
	DBPort        string
	DBName        string
	DBSchema      string
	DBSSLMode     string
	DBConnTimeout int64
}

type Config struct {
	PublicHost             string
	ServerPort             string
	JWTSecret              string
	JWTExpirationInSeconds int64
	PGConfig               *PGConfig
}

func InitConfig() Config {
	pgconfig := PGConfig{
		DBHost:        getEnvStr("ATLAS_HOST", "localhost"),
		DBPort:        getEnvStr("ATLAS_PORT", "5436"),
		DBUser:        getEnvStr("LOSLOGGER_DB_USER", ""),
		DBPassword:    getEnvStr("LOSLOGGER_DB_PASS", ""),
		DBName:        getEnvStr("LOSLOGGER_DB_NAME", "loslogger"),
		DBSchema:      getEnvStr("LOSLOGGER_DB_SCHEMA", "loslogger"),
		DBSSLMode:     getEnvStr("LOSLOGGER_DB_SSLMODE", "require"),
		DBConnTimeout: getEnvInt64("LOSLOGGER_DB_CONNTIMEOUT", 0),
	}

	return Config{
		PublicHost:             getEnvStr("LOSLOGGER_PUBLIC_HOST", ""),
		ServerPort:             getEnvStr("LOSLOGGER_API_PORT", "9031"),
		JWTSecret:              getEnvStr("LOSLOGGER_JWT_SECRET", ""),
		JWTExpirationInSeconds: getEnvInt64("LOSLOGGER_JWT_EXPIRATION_IN_SECONDS", 3600*24*7),
		PGConfig:               &pgconfig,
	}
}

func getEnvStr(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvInt64(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}
