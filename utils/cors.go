package utils

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func DefaultCors() cors.Config {
	config := cors.ConfigDefault
	config.AllowOrigins = "*"
	config.AllowHeaders = "Content-Type,access-control-allow-origin, access-control-allow-headers, token, id, companyid"
	// config.AllowMethods = "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS"
	// config.AllowCredentials = true
	return config
}
