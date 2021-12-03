package main

import (
	"fmt"
	"github.com/jhonyzam/curso_golang/aula_8/products/domain/product"
	"github.com/jhonyzam/curso_golang/aula_8/products/internal/infrastructure/server/http"
	"github.com/jhonyzam/curso_golang/aula_8/products/internal/infrastructure/storage/postgres"
	"github.com/jhonyzam/curso_golang/aula_8/products/pkg/env"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	envApplicationPort = "PRODUCTS_PORT"
	envPostgresURL     = "POSTGRES_URL"
	defaultPort        = "8080"
)

var (
	version, date string
)

func main() {
	log.Printf("Products API starting - version: %s; date: %s\n", version, date)

	//env.CheckRequired(envPostgresURL)

	// Storages...
	db, err := postgres.NewConnection(getPostgresURL())
	if err != nil {
		log.Fatalf("ERROR: connecting database: %q\n", err)
		return
	}

	productStorage := postgres.NewProductStorage(db)

	/*
	* Services...
	 */
	productService := product.NewService(productStorage)

	/*
	* Handler...
	 */
	handler := http.NewHandler(productService)

	/*
	* Server...
	 */
	server := http.New(getApplicationPort(), handler)
	server.ListenAndServe()

	/*
	* Graceful shutdown...
	 */
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)
	<-stopChan
	server.Shutdown()
}

func getApplicationPort() string {
	return env.GetString(envApplicationPort, defaultPort)
}

func getPostgresURL() string {
	fmt.Println(env.GetString(envPostgresURL))
	return env.GetString(envPostgresURL)
}
