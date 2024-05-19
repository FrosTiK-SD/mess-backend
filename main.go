package main

import (
	"fmt"
	"os"

	"github.com/FrosTiK-SD/mess-backend/constants"
	"github.com/FrosTiK-SD/mess-backend/handler"
	"github.com/FrosTiK-SD/mess-backend/interfaces"
	"github.com/FrosTiK-SD/mess-backend/utils"
	"github.com/FrosTiK-SD/mongik"
	mongikConstants "github.com/FrosTiK-SD/mongik/constants"
	mongikConfig "github.com/FrosTiK-SD/mongik/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/swagger"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func main() {

	app := fiber.New(fiber.Config{
		Prefork:           true,
		JSONEncoder:       json.Marshal,
		EnablePrintRoutes: true,
		JSONDecoder:       json.Unmarshal,
		// Global custom error handler
		ErrorHandler: func(c *fiber.Ctx, err error) error {

			return c.Status(fiber.StatusBadRequest).JSON(interfaces.GetGenericResponse(false, "ERROR", nil, err))
		},
	})

	mongikClient := mongik.NewClient(os.Getenv(constants.CONNECTION_STRING), &mongikConfig.Config{
		Client: mongikConstants.BIGCACHE,
		TTL:    constants.CACHING_DURATION,
		Debug:  false,
		RedisConfig: &mongikConfig.RedisConfig{
			URI:      os.Getenv(constants.REDIS_URI),
			Password: os.Getenv(constants.REDIS_PASSWORD),
			Username: os.Getenv(constants.REDIS_USERNAME),
		},
		FallbackToDefault: true,
	})

	handler := &handler.Handler{
		MongikClient: mongikClient,
	}

	// Allow origin
	app.Use(cors.New(utils.DefaultCors()))

	// Recover from panics || Comment this out to check panic logs
	// app.Use(recover.New())

	// Rate limiting
	app.Use(limiter.New(limiter.Config{Max: constants.REQUEST_RATE}))

	// Compress responses
	app.Use(compress.New())

	// Security
	app.Use(helmet.New())

	// Health check
	app.Use(healthcheck.New())

	app.Get("/hello", handler.RegisterStudent)

	// Monitor
	app.Get("/metrics", monitor.New())
	app.Get("/swagger/*", swagger.HandlerDefault)

	port := "" + os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Starting Server on PORT : ", port)

	app.Listen(":" + port)
}
