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

	if err := godotenv.Load(); err != nil {
		logger.Warn(context.Background(), "[server] No se encontró archivo .env, se usarán variables del sistema")
	}

	db, err := providers.DatabaseConnectionPostgres()
	if err != nil {
		logger.Error(context.Background(), "[server] Error al conectar a la base de datos: %v", err)
		os.Exit(1)
	}
	middlewares2.InitRoleMiddleware(db)

	appInstance, err := app.Start()
	if err != nil {
		logger.Error(context.Background(), "[server] Error al inicializar la aplicación: %v", err)
		os.Exit(1)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = os.Getenv("APP_PORT")
		if port == "" {
			port = "8080"
		}
	}

	addr := ":" + port

	logger.Info(context.Background(), fmt.Sprintf("[server] Configurando servidor en el puerto %s", addr))

	go func() {
		logger.Info(context.Background(), "[server] Iniciando servidor HTTP...")
		if err := appInstance.Start(addr); err != nil {
			logger.Error(context.Background(), "[server] Error al iniciar el servidor: %v", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info(context.Background(), "[server] Señal de apagado recibida, cerrando servidor...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := appInstance.Shutdown(ctx); err != nil {
		logger.Error(context.Background(), "[server] Error al cerrar el servidor: %v", err)
	}

	logger.Info(ctx, "[server] Servidor cerrado correctamente ✅")
}
