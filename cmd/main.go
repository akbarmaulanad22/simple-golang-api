package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"api/config"
	"api/internal/interface/controller"
	"api/internal/interface/middleware"
	"api/internal/interface/route"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize DB
	db := config.InitDB()

	// Initialize router
	router := mux.NewRouter()

	// Add middleware
	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.RecoveryMiddleware)

	// === PUBLIC ROUTES (tanpa middleware) ===
	authController := controller.NewAuthController(db)
	
    publicRoutes := router.PathPrefix("/").Subrouter()
    publicRoutes.HandleFunc("/login", authController.Login).Methods("POST")

	// === PROTECTED ROUTES (dengan middleware auth) ===
    protectedRoutes := router.PathPrefix("/api/v1").Subrouter()
    protectedRoutes.Use(middleware.AuthMiddleware)
	
	// Setup routes
	route.SetupRoutes(protectedRoutes, db)

	// Konfigurasi server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Server dijalankan dalam goroutine terpisah
	go func() {
		fmt.Printf("=========================================\n")
		fmt.Printf("Server berjalan pada port %s...\n", port)
		fmt.Printf("Endpoint yang tersedia:\n")
		fmt.Printf("http://localhost:%s/api/v1/user\n", port)

		fmt.Printf("=========================================\n")

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error menjalankan server: %s\n", err)
		}
	}()

	// Channel untuk menangkap signal shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %s\n", err)
	}

	log.Println("Server berhasil shutdown")
}
