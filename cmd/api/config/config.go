package config

import (
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

	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No se pudo cargar el archivo .env, usando variables de entorno del sistema.")
	}

	env := os.Getenv("GO_ENVIRONMENT")
	if env == "" {
		env = constants.ScopeLocal
	}

	dbConfig := ConnectionConfig{
		Username:             getEnv("DB_USER", "postgres"),
		Password:             getEnv("DB_PASSWORD", "Pilotof1988*"),
		Host:                 getEnv("DB_HOST", "localhost"),
		Name:                 getEnv("DB_NAME", "postgres"),
		Port:                 getEnv("DB_PORT", "5432"),
		SSLMode:              getEnv("DB_SSLMODE", "require"),
		MaxIdleConnections:   500,
		MaxOpenConnections:   500,
		MaxConnectionRetries: 3,
		MaxBatchSize:         100,
		ConnMaxLifetime:      600 * time.Second,
		ConnMaxIdleTime:      600 * time.Second,
	}

	jwtConfig := JWTManager{
		SecretKey:     getEnv("JWT_SECRET", "your-256-bit-secret"),
		AccessExpiry:  15 * time.Minute,
		RefreshExpiry: 90 * 24 * time.Hour,
		ExpiredTokens: 24 * time.Hour,
	}

	rustyConfig := RustyClientConfig{
		DefaultTimeOut: 11 * time.Second,
		RetryCount:     3,
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
