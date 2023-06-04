package router

import (
	"../middleware"
	"github.com/gofiber/fiber/v2"
)

// Router is exported and used in main.go
func Router() *fiber.App {
	router := fiber.New()

	router.Get("/api/task", middleware.GetAllTask)
	router.Post("/api/task", middleware.GetAllTask)
	router.Put("/api/task/:id", middleware.GetAllTask)
	router.Delete("/api/deleteTask/:id", middleware.GetAllTask)

	return router
}
