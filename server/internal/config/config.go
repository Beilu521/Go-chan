package config

import (
	"log/slog"
	"os"
	"strconv"
	"github.com/joho/godotenv"
)

type Config struct{
	AppEnv string
	AppPort string
	DatabaseDsn string
	JwtSecret string
	JwtExpireHours int
	EncryptionKey string
	AgentMode string
	HermesBaseUrl string
}

func Load() Config{
	err := godotenv.Load()
	if err != nil{
		slog.Error("failed to load .env file","err",err)
		os.Exit(1)
	}
	
	expireHours,err := strconv.Atoi(getEnv("JWT_EXPIRE_HOURS","720"))
	if err != nil{
		slog.Warn("failed to parse JWT_EXPIRE_HOURS, using default 720")
		expireHours = 720
	}

	cfg := Config{
		AppEnv: getEnv("APP_ENV","dev"),
		AppPort: getEnv("APP_PORT","8080"),
		DatabaseDsn: getEnv("DATABASE_DSN",""),
		JwtSecret: getEnv("JWT_SECRET",""),
		JwtExpireHours: expireHours,
		EncryptionKey: getEnv("APP_ENCRYPTION_KEY",""),
		AgentMode: getEnv("AGENT_MODE","mock"),
		HermesBaseUrl: getEnv("HERMES_BASE_URL","http://localhost:9000"),
	}

	if cfg.DatabaseDsn == ""{
		slog.Error("DATABASE_DSN is empty")
		os.Exit(1)
	}
	if cfg.JwtSecret == ""{
		slog.Error("JWT_SECRET is empty")
		os.Exit(1)
	}
	if cfg.EncryptionKey == ""{
		slog.Error("APP_ENCRYPTION_KEY is empty")
		os.Exit(1)
	}
	if len(cfg.EncryptionKey) != 32{
		slog.Error("APP_ENCRYPTION_KEY is invalid (must be 32 chars)", "len", len(cfg.EncryptionKey))
		os.Exit(1)
	}

	slog.Info("config loaded successfully")
	return cfg
}

func getEnv(key string,fallback string) string{
	value := os.Getenv(key)
	if value == ""{
		return fallback
	}
	return value
}