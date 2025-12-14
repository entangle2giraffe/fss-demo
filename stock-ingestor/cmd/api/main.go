package main

import (
	"fmt"
	"log"
	"time"
	"stock-ingestor/config"
	"stock-ingestor/internal/application/usecases"
	alphavantage "stock-ingestor/internal/infrastructure/external/alphavantage"
	"stock-ingestor/internal/interfaces/http/handlers"
	"stock-ingestor/internal/interfaces/http/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize repository
	repo := alphavantage.NewStockRepositoryImpl(cfg.AlphaVantage.APIKey, cfg.AlphaVantage.URI)

	// Initialize use case
	getStockTimeSeriesUseCase := usecases.NewGetStockTimeSeriesUseCase(repo)
	getCompanyOverviewUseCase := usecases.NewGetCompanyOverviewUseCase(repo)

	// Initialize handler
	stockHandler := handlers.NewStockHandler(getStockTimeSeriesUseCase, getCompanyOverviewUseCase)

	// Initialize Fiber app
	app := fiber.New()

	app.Use(healthcheck.New())
	app.Use(csrf.New())
	app.Use(logger.New())

	// Add CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept, X-API-Key",
	}))

	// Add cache middleware for GET requests
	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Method() != "GET"
		},
		Expiration: 5 * time.Minute, // 5 minutes
	}))

	// Routes
	app.Get("/api/stocks/:symbol", middleware.APIKeyAuth(cfg.Server.APIKey), stockHandler.GetStockTimeSeries)
	app.Get("/api/stocks/:symbol/overview", middleware.APIKeyAuth(cfg.Server.APIKey), stockHandler.GetCompanyOverview)

	log.Fatal(app.Listen(":" + fmt.Sprintf("%d", cfg.Server.Port)))
}
