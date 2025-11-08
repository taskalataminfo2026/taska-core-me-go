package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/taskalataminfo2026/tool-kit-lib-go/pkg/logger"
	"os"
	"os/signal"
	"syscall"
	"taska-core-me-go/cmd/api/app"
	"taska-core-me-go/cmd/api/app/providers"
	"taska-core-me-go/cmd/api/constants"
	middlewares2 "taska-core-me-go/cmd/api/middlewares"
	"time"
)

// @title Taska Auth API
// @version 1.0
// @description Servicio Core
// @termsOfService https://taska/terminos

// @contact.name Soporte Taska
// @contact.email taska.latam.info@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /v1/api/core
func main() {
	logger.Init()
	defer logger.Sync()

	ctx := context.Background()
	logger.Info(ctx, "[server] Inicializando aplicación...")

	// 🔹 Determinar entorno y archivo .env correspondiente
	env := os.Getenv("GO_ENVIRONMENT")
	if env == "" {
		env = constants.ScopeLocal // por defecto "local"
	}

	envFile := ".env.local"
	switch env {
	case constants.ScopeTest:
		envFile = ".env.test"
	case constants.ScopeProduction:
		envFile = ".env.prod"
	}

	// 🔹 Cargar variables de entorno
	if err := godotenv.Load(envFile); err != nil {
		logger.Warn(ctx, fmt.Sprintf("[server] No se encontró %s, se usarán variables del sistema", envFile))
	} else {
		logger.Info(ctx, fmt.Sprintf("[server] Variables de entorno cargadas desde %s ✅", envFile))
	}

	// 🔹 Conexión a base de datos
	db, err := providers.DatabaseConnectionPostgres()
	if err != nil {
		logger.Error(ctx, "[server] Error al conectar a la base de datos: %v", err)
		os.Exit(1)
	}
	middlewares2.InitRoleMiddleware(db)

	// 🔹 Inicializar la app
	appInstance, err := app.Start()
	if err != nil {
		logger.Error(ctx, "[server] Error al inicializar la aplicación: %v", err)
		os.Exit(1)
	}

	// 🔹 Puerto de ejecución
	port := os.Getenv("PORT")
	if port == "" {
		port = os.Getenv("APP_PORT")
		if port == "" {
			port = "8080"
		}
	}
	addr := ":" + port
	logger.Info(ctx, fmt.Sprintf("[server] Configurando servidor en el puerto %s", addr))

	// 🔹 Ejecutar servidor
	go func() {
		logger.Info(ctx, "[server] Iniciando servidor HTTP...")
		if err := appInstance.Start(addr); err != nil {
			logger.Error(ctx, "[server] Error al iniciar el servidor: %v", err)
			os.Exit(1)
		}
	}()

	// 🔹 Capturar señal de apagado
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info(ctx, "[server] Señal de apagado recibida, cerrando servidor...")

	// 🔹 Cierre limpio del servidor
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := appInstance.Shutdown(shutdownCtx); err != nil {
		logger.Error(ctx, "[server] Error al cerrar el servidor: %v", err)
	}

	logger.Info(ctx, "[server] Servidor cerrado correctamente ✅")
}
