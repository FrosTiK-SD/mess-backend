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
	"github.com/joho/godotenv"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func main() {

	godotenv.Load()

	app := fiber.New(fiber.Config{
		Prefork:           false,
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

	app.Get("/hello", handler.Hello)

	adminAPI := app.Group("/admin")
	{

		adminAPI.Post("/mess", handler.CreateMess)
		adminAPI.Get("/mess", handler.GetMess)

		adminAPI.Post("/hostel", handler.CreateHostel)
		adminAPI.Get("/hostel", handler.GetHostel)

		adminAPI.Post("/meal", handler.CreateMeal)
		adminAPI.Get("/meal", handler.GetMeal)

		adminAPI.Post("/meal-type", handler.CreateMealType)
		adminAPI.Get("/meal-type", handler.GetMealType)

		adminAPI.Post("/menu-item", handler.CreateMenuItem)
		adminAPI.Get("/menu-item", handler.GetMenuItem)

		adminAPI.Post("/room", handler.CreateRoom)

		adminAPI.Post("/user", handler.CreateUser)
		adminAPI.Get("/user/populated", handler.GetUserPopulated)
		adminAPI.Post("/user/manage/hostel-mess", handler.ManageHostelMess)

		adminAPI.Get("/filteredStudents", handler.GetFilteredStudents)
	}

	caretakerAPI := app.Group("/caretaker")
	{
		caretakerAPI.Get("/mess/dashboard", handler.GetMessDashboard)
		caretakerAPI.Get("/menu-item/all", handler.GetAllMenuItemsOfAMess)
		caretakerAPI.Get("/meal-type/all", handler.GetAllMealTypesOfAMess)
		caretakerAPI.Post("/meal/update/menu/by/date", handler.UpdateMealsByDate)
	}

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
