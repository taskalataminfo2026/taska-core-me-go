package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/taskalataminfo2026/tool-kit-lib-go/pkg/logger"
	"taska-core-me-go/cmd/api/app"
)

// @title Taska Auth API
// @version 1.0
// @description API de autenticación y gestión de usuarios para Taska LATAM.
// @termsOfService https://taska/terminos

// @contact.name Soporte Taska
// @contact.email taska.latam.info@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /v1/api/core
func main() {
	ctx := context.Background()

	logger.Init()
	defer logger.Sync()

	logger.Info(ctx, "[server] Inicializando aplicación...")

	appInstance, err := app.Start()
	if err != nil {
		logger.Error(ctx, "[server] Error al inicializar la aplicación: %v", err)
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
	logger.Info(ctx, fmt.Sprintf("[server] Configurando servidor en el puerto %s", addr))

	go func() {
		logger.Info(ctx, "[server] Iniciando servidor HTTP...")
		if err := appInstance.Start(addr); err != nil {
			logger.Error(ctx, "[server] Error al iniciar el servidor: %v", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info(ctx, "[server] Señal de apagado recibida, cerrando servidor...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := appInstance.Shutdown(shutdownCtx); err != nil {
		logger.Error(ctx, "[server] Error al cerrar el servidor: %v", err)
	}

	logger.Info(ctx, "[server] Servidor cerrado correctamente")
}
