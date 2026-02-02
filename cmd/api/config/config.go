package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
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

type RustyClientConfig struct {
	DefaultTimeOut time.Duration
	RetryCount     int
}

type AppConfig struct {
	DB    ConnectionConfig
	Rusty RustyClientConfig
	Env   string
}

var Config *AppConfig

func init() {
	env := os.Getenv("GO_ENVIRONMENT")
	if env == "" {
		env = constants.ScopeLocal
	}

	loadEnvFile(env)

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

	rustyConfig := RustyClientConfig{
		DefaultTimeOut: getDurationEnv("RUSTY_TIMEOUT", 11) * time.Second,
		RetryCount:     getIntEnv("RUSTY_RETRY_COUNT", 3),
	}

	Config = &AppConfig{
		DB:    dbConfig,
		Rusty: rustyConfig,
		Env:   env,
	}
}

func loadEnvFile(env string) {
	envFiles := map[string]string{
		constants.ScopeLocal:      ".env.local",
		constants.ScopeTest:       ".env.test",
		constants.ScopeProduction: ".env.prod",
	}

	envFile, ok := envFiles[env]
	if !ok {
		log.Fatalf("No hay archivo de entorno configurado para %s", env)
	}

	projectRoot := findProjectRoot()
	fullPath := filepath.Join(projectRoot, envFile)

	if err := godotenv.Load(fullPath); err != nil {
		log.Printf("No se pudo cargar %s, usando variables del sistema", fullPath)
	} else {
		log.Printf("Cargado archivo de entorno: %s", fullPath)
	}
}

// findProjectRoot busca la raíz del proyecto buscando go.mod
func findProjectRoot() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	log.Fatal("No se pudo encontrar el proyecto raíz (go.mod)")
	return ""
}

// getEnv obtiene una variable de entorno o devuelve fallback
func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

// getIntEnv obtiene una variable de entorno como int o devuelve fallback
func getIntEnv(key string, fallback int) int {
	if val := os.Getenv(key); val != "" {
		var i int
		if _, err := fmt.Sscanf(val, "%d", &i); err == nil {
			return i
		}
	}
	return fallback
}

// getDurationEnv obtiene una variable de entorno como duración en segundos o devuelve fallback
func getDurationEnv(key string, fallback int) time.Duration {
	if val := os.Getenv(key); val != "" {
		var i int
		if _, err := fmt.Sscanf(val, "%d", &i); err == nil {
			return time.Duration(i)
		}
	}
	return time.Duration(fallback)
}
