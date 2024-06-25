package config

import (
	"log"
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
	LogFileLoc string
	WebBeacon  string
	PGConfig   *PGConfig
}

func InitConfig() Config {
	pgconfig := PGConfig{
		DBHost:        getEnvStr("ATLAS_HOST", "localhost", true),
		DBPort:        getEnvStr("ATLAS_PORT", "5436", true),
		DBUser:        getEnvStr("LOSL_DB_USER", "", true),
		DBPassword:    getEnvStr("LOSL_DB_PASS", "", true),
		DBName:        getEnvStr("LOSL_DB_NAME", "loslogger", true),
		DBSchema:      getEnvStr("LOSL_DB_SCHEMA", "loslogger", true),
		DBSSLMode:     getEnvStr("LOSL_DB_SSLMODE", "require", true),
		DBConnTimeout: getEnvInt64("LOSL_DB_CONNTIMEOUT", 0, true),
	}

	return Config{
		LogFileLoc: getEnvStr("LOSL_LOGFILELOC", "", true),
		WebBeacon:  getEnvStr("LOSL_WBEACON", "", true),
		PGConfig:   &pgconfig,
	}
}

func getEnvStr(key, fallback string, req bool) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	if req {
		log.Fatalf("Env key, %s, does not have a value", key)
	}
	return fallback
}

func getEnvInt64(key string, fallback int64, req bool) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			if req {
				log.Fatalf("Env key, %s, does not have a value", key)
			}
			return fallback
		}
		return i
	}
	if req {
		log.Fatalf("Env key, %s, does not have a value", key)
	}
	return fallback
}
