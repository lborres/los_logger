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
	LogFileLoc             string
	WebBeacon              string
	PGConfig               *PGConfig
}

func InitConfig() Config {
	pgconfig := PGConfig{
		DBHost:        getEnvStr("ATLAS_HOST", "localhost"),
		DBPort:        getEnvStr("ATLAS_PORT", "5436"),
		DBUser:        getEnvStr("LOSL_DB_USER", ""),
		DBPassword:    getEnvStr("LOSL_DB_PASS", ""),
		DBName:        getEnvStr("LOSL_DB_NAME", "loslogger"),
		DBSchema:      getEnvStr("LOSL_DB_SCHEMA", "loslogger"),
		DBSSLMode:     getEnvStr("LOSL_DB_SSLMODE", "require"),
		DBConnTimeout: getEnvInt64("LOSL_DB_CONNTIMEOUT", 0),
	}

	return Config{
		PublicHost:             getEnvStr("LOSL_PUBLIC_HOST", ""),
		ServerPort:             getEnvStr("LOSL_API_PORT", "9031"),
		JWTSecret:              getEnvStr("LOSL_JWT_SECRET", ""),
		JWTExpirationInSeconds: getEnvInt64("LOSL_JWT_EXPIRATION_IN_SECONDS", 3600*24*7),
		LogFileLoc:             getEnvStr("LOSL_LOGFILE", "/app/los_logger.log"),
		WebBeacon:              getEnvStr("LOSL_WBEACON", ""),
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
