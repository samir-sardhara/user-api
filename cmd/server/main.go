package main

import (
	"database/sql"
	"log"
	"user-api/config"
	"user-api/internal/handler"
	"user-api/internal/logger"
	"user-api/internal/middleware"
	"user-api/internal/repository"
	"user-api/internal/routes"
	"user-api/internal/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	_ "github.com/gofiber/fiber/v2/middleware/requestid"
	"go.uber.org/zap"
)

func main() {
	// Initialize global logging matrix
	logger.InitLogger()
	defer logger.Log.Sync()

	// Parse internal configuration parameters
	cfg := config.LoadConfig()

	// Initialize MySQL network connection matrix
	db, err := sql.Open("mysql", cfg.DBUrl)
	if err != nil {
		logger.Log.Fatal("System state creation aborted. Connection configuration failure", zap.Error(err))
	}
	defer db.Close()

	// Ensure structural persistence target is reachable
	if err := db.Ping(); err != nil {
		logger.Log.Fatal("System state creation aborted. Target system instance unreachable", zap.Error(err))
	}
	logger.Log.Info("System layer established. Connection vector actively handling structural requests")

	// Inject structures through application architecture layers
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Instantiate operational web layout context
	app := fiber.New()
	app.Use(requestid.New())
	app.Use(middleware.ZapLogger())

	// Build pathway layout context links
	routes.SetupRoutes(app, userHandler)

	// Open local network system socket
	logger.Log.Info("Executing interface engine processing on port configuration mapping " + cfg.Port)
	if err := app.Listen(cfg.Port); err != nil {
		log.Fatalf("Fatal network lifecycle loop failure triggered: %v", err)
	}
}
