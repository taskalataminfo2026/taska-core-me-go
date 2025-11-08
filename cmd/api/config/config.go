package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"taska-core-me-go/cmd/api/constants"
	"time"
)

type ConnectionConfig struct {
	Username             string
	Password             string
	Host                 string
	Name                 string
	Port                 string
	SSLMode              string
	MaxIdleConnections   int
	MaxOpenConnections   int
	MaxConnectionRetries int
	MaxBatchSize         int
	ConnMaxLifetime      time.Duration
	ConnMaxIdleTime      time.Duration
}

type JWTManager struct {
	SecretKey     string
	AccessExpiry  time.Duration
	RefreshExpiry time.Duration
	ExpiredTokens time.Duration
}

type RustyClientConfig struct {
	DefaultTimeOut time.Duration
	RetryCount     int
}

type AppConfig struct {
	DB    ConnectionConfig
	JWT   JWTManager
	Rusty RustyClientConfig
	Env   string
}

var Config *AppConfig

func init() {
	env := os.Getenv("GO_ENVIRONMENT")
	if env == "" {
		env = constants.ScopeLocal
	}

	// Selecciona el archivo de entorno correcto según el ambiente
	envFile := ".env.local"
	switch env {
	case constants.ScopeTest:
		envFile = ".env.test"
	case constants.ScopeProduction:
		envFile = ".env.prod"
	}

	if err := godotenv.Load(envFile); err != nil {
		log.Printf("⚠️  No se pudo cargar %s, usando variables del sistema.", envFile)
	}

	dbConfig := ConnectionConfig{
		Username:             getEnv("DB_USER", "postgres"),
		Password:             getEnv("DB_PASSWORD", ""),
		Host:                 getEnv("DB_HOST", "localhost"),
		Name:                 getEnv("DB_NAME", "postgres"),
		Port:                 getEnv("DB_PORT", "5432"),
		SSLMode:              getEnv("DB_SSLMODE", "disable"),
		MaxIdleConnections:   50,
		MaxOpenConnections:   100,
		MaxConnectionRetries: 3,
		MaxBatchSize:         100,
		ConnMaxLifetime:      600 * time.Second,
		ConnMaxIdleTime:      600 * time.Second,
	}

	jwtConfig := JWTManager{
		SecretKey:     getEnv("JWT_SECRET", "default-secret"),
		AccessExpiry:  15 * time.Minute,
		RefreshExpiry: 90 * 24 * time.Hour,
		ExpiredTokens: 24 * time.Hour,
	}

	rustyConfig := RustyClientConfig{
		DefaultTimeOut: getDurationEnv("RUSTY_TIMEOUT", 11) * time.Second,
		RetryCount:     getIntEnv("RUSTY_RETRY_COUNT", 3),
	}

	Config = &AppConfig{
		DB:    dbConfig,
		JWT:   jwtConfig,
		Rusty: rustyConfig,
		Env:   env,
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func getIntEnv(key string, fallback int) int {
	if val := os.Getenv(key); val != "" {
		var i int
		if _, err := fmt.Sscanf(val, "%d", &i); err == nil {
			return i
		}
	}
	return fallback
}

func getDurationEnv(key string, fallback int) time.Duration {
	if val := os.Getenv(key); val != "" {
		var i int
		if _, err := fmt.Sscanf(val, "%d", &i); err == nil {
			return time.Duration(i)
		}
	}
	return time.Duration(fallback)
}
