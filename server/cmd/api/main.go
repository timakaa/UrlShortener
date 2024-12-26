package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/timakaa/test-go/internal/db"
	"github.com/timakaa/test-go/internal/handlers"
	"github.com/timakaa/test-go/internal/middleware"
	"github.com/timakaa/test-go/internal/server"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(fmt.Sprintf("Warning: Error loading .env file: %v\n", err))
	}
	
	// Initialize database
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}

	if err := db.InitRedis(); err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	srv := server.New()
	
	srv.Use(
		middleware.CorsMiddleware(),
		middleware.SecurityMiddleware(),
	)

	srv.GET("/api/hello", handlers.HelloHandler, middleware.AuthMiddleware())
	srv.POST("/api/login", handlers.LoginHandler)
	srv.POST("/api/register/init", handlers.RegisterInitHandler)
	srv.POST("/api/register/verify", handlers.VerifyAndRegisterHandler)
	srv.POST("/api/logout", handlers.LogoutHandler)
	srv.POST("/api/url", handlers.CreateUrlHandler, middleware.AuthMiddleware())
	srv.GET("/api/url", handlers.GetUrlHandler)
	srv.GET("/api/urls", handlers.GetUrlsHandler, middleware.AuthMiddleware())

	srv.ListenAndServe()
}