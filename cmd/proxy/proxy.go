package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"proxy_server/internal/config"
	"proxy_server/internal/logging"
	"proxy_server/internal/proxy"
	"proxy_server/internal/proxy/handlers"
	proxytls "proxy_server/internal/proxy/tls"
	"proxy_server/internal/service"
	"proxy_server/internal/storage"
	"proxy_server/internal/storage/postgresql"

	"github.com/asaskevich/govalidator"
)

const configPath string = "config/config.yml"
const envPath string = "config/.env"

func main() {
	config, err := config.LoadConfig(envPath, configPath)
	govalidator.SetFieldsRequiredByDefault(true)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("Config loaded")

	logger, err := logging.NewLogrusLogger(config.Logging)
	if err != nil {
		logger.Fatal(err.Error())
	}
	logger.Info("Logger configured")

	ca, err := proxytls.LoadCA(config.TLS, &logger)
	if err != nil {
		logger.Fatal(err.Error())
	}
	logger.Info("TLS CA loaded")

	dbConnection, err := postgresql.GetDBConnection(*config.Database)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer dbConnection.Close()
	logger.Info("Database connection established")

	storages := storage.NewPostgresStorages(dbConnection)
	logger.Info("Storages configured")

	services := service.NewServices(storages)
	logger.Info("Services configured")

	handlers := handlers.NewHandlers(services, ca, config.TLS)
	logger.Info("Handlers configured")

	mux, err := proxy.GetChiMux(*handlers, *config, &logger)
	if err != nil {
		logger.Fatal(err.Error())
	}
	logger.Info("Router configured")

	var server = http.Server{
		Addr:    fmt.Sprintf(":%d", config.Proxy.Port),
		Handler: mux,
	}

	logger.Info("Server is running")

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// We received an interrupt signal, shut down.
		if err := server.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			logger.Info("HTTP server Shutdown: " + err.Error())
		}
		close(idleConnsClosed)
	}()

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		logger.Fatal("HTTP server ListenAndServe: " + err.Error())
	}

	<-idleConnsClosed
}
