package router

import (
	"todo-list-app/server/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	_ "github.com/mdzkm/todo-list-app/server/docs"
)

// Router is exported and used in main.go
func Router() *fiber.App {
	router := fiber.New()

	// Middleware
	router.Use(recover.New())
	router.Use(cors.New())
 
	// Routes
	router.Get("/", middleware.HealthCheck)
	router.Get("/api/tasks", middleware.GetTasks)
	router.Post("/api/tasks", middleware.CreateTask)
	router.Put("/api/tasks/update/:id", middleware.UpdateTask)
	router.Put("/api/tasks/complete/:id", middleware.CompleteTask)
	router.Delete("/api/tasks/delete/:id", middleware.DeleteTask)
	router.Get("/swagger/*", swagger.HandlerDefault)

	return router
}
