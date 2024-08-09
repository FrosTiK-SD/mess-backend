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
		adminAPI.Put("/mess", handler.UpdateMess)
		adminAPI.Delete("/mess", handler.DeleteMess)

		adminAPI.Post("/hostels", handler.CreateHostel)
		adminAPI.Get("/hostels", handler.GetAllHostels)
		adminAPI.Get("/hostels/:hostelId", handler.GetHostelById)
		adminAPI.Get("/populatedHostels/:hostelId", handler.GetFullyPopulatedHostel)
		adminAPI.Put("/hostels", handler.UpdateHostel)
		adminAPI.Delete("/hostels", handler.DeleteHostel)

		adminAPI.Post("/meal", handler.CreateMeal)
		adminAPI.Get("/meal", handler.GetMeal)
		adminAPI.Put("/meal", handler.UpdateMeal)
		adminAPI.Delete("/meal", handler.DeleteMeal)

		adminAPI.Post("/meal-type", handler.CreateMealType)
		adminAPI.Get("/meal-type", handler.GetMealType)
		adminAPI.Put("/meal-type", handler.UpdateMealType)
		adminAPI.Delete("/meal-type", handler.DeleteMealType)

		adminAPI.Post("/menu-item", handler.CreateMenuItem)
		adminAPI.Get("/menu-item", handler.GetMenuItem)
		adminAPI.Put("/menu-item", handler.UpdateMenuItem)
		adminAPI.Delete("/menu-item", handler.DeleteMenuItem)

		adminAPI.Post("/room", handler.CreateRoom)

		adminAPI.Post("/user", handler.CreateUser)
		adminAPI.Get("/user", handler.GetUser)
		adminAPI.Put("/user", handler.UpdateUser)
		adminAPI.Delete("/user", handler.DeleteUser)

		adminAPI.Get("/user/populated", handler.GetUserPopulated)
		adminAPI.Get("/user/rollNo/:rollNo", handler.GetUserByRollNo)
		adminAPI.Post("/user/manage/hostel-mess", handler.ManageHostelMess)

		adminAPI.Post("/userFiltered", handler.GetFilteredUsers)
		adminAPI.Put("/user/assignHostel", handler.AssignHostelToUsers)
		adminAPI.Put("/user/assignMess", handler.AssignMessToUsers)
		adminAPI.Put("/user/assignRoom", handler.AssignRoomToUser)

		adminRooms := adminAPI.Group("/rooms")
		{
			adminRooms.Post("/generate", handler.GenerateRooms)
		}
	}

	caretakerAPI := app.Group("/caretaker")
	{
		caretakerAPI.Get("/mess/dashboard", handler.GetMessDashboard)
		caretakerAPI.Get("/menu-item/all", handler.GetAllMenuItemsOfAMess)
		caretakerAPI.Get("/meal-type/all", handler.GetAllMealTypesOfAMess)
		caretakerAPI.Post("/meal/update/menu/by/date", handler.UpdateMealsByDate)
	}

	userAPI := app.Group("/user")
	{
		userAPI.Get("/token", handler.FiberAuthenticateUser, handler.GetUserFromToken)
		userAPI.Post("/token", handler.CreateUserFromToken)
	}

	// Monitor
	app.Get("/metrics", monitor.New())
	app.Get("/swagger/*", swagger.HandlerDefault)

	port := "" + os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	port = ":" + port
	if os.Getenv("APP_ENV") == "dev" {
		port = "localhost" + port
	}

	fmt.Println("Starting Server on PORT : ", port)

	app.Listen(port)
}
