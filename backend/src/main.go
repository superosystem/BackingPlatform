package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/superosystem/BackingPlatform/backend/src/application"
	"github.com/superosystem/BackingPlatform/backend/src/config"
	"github.com/superosystem/BackingPlatform/backend/src/middleware"
)

func main() {
	// SETUP ENVIRONMENT
	cfg := config.NewConfig()
	dbConnect := config.NewDBConn(cfg)

	PORT := cfg.App.Port
	if PORT == "" {
		PORT = "8080"
	}

	// SETUP SERVER
	gin.SetMode(cfg.App.Mode)
	server := gin.Default()
	server.Use(cors.Default())
	server.Use(func(ctx *gin.Context) {
		// ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}

	})

	cookieStore := cookie.NewStore([]byte(middleware.SECRET_KEY))
	server.Use(sessions.Sessions("backing-platform", cookieStore))
	// FILE STORAGE
	server.Static("/public", "./public")

	// DEPENDENCY INJECTION
	service := application.NewService(application.NewRepository(dbConnect))
	// API
	application.StartApiV1(server, service)
	// Web
	application.StartWeb(server, service)

	// STARTING...
	server.Run(PORT)
}
