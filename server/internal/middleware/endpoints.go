package middleware

import (
	"github.com/mdzkm/todo-list-app/server/internal/controllers"
	"github.com/mdzkm/todo-list-app/server/internal/models"

	"github.com/gofiber/fiber/v2"
)

// GetTasks godoc
// @Summary Show tasks
// @Description get all tasks
// @Tags tasks
// @Produce json
func GetTasks(c *fiber.Ctx) error {
	payload := controllers.GetTasks()
	c.JSON(payload)
	return nil
}

// CreateTasks godoc
// @Summary Create a task
// @Description create a single task
// @Tags tasks
// @Accept json
// @Produce json
func CreateTask(c *fiber.Ctx) error {
	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return err
	}
	controllers.CreateTask(task)
	c.JSON(task)
	return nil
}

// UpdateTasks godoc
// @Summary Update task description
// @Description replace task description by ID
// @Tags tasks
// @Accept json
// @Produce json
func UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return err
	}
	controllers.UpdateTask(id, task.Description)
	c.JSON(id)
	return nil
}

// CompleteTasks godoc
// @Summary Complete a task
// @Description change status of task to complete by ID
// @Tags tasks
// @Accept json
// @Produce json
func CompleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	controllers.CompleteTask(id)
	c.JSON(id)
	return nil
}

// DeleteTasks godoc
// @Summary Delete a task
// @Description delete task by ID
// @Tags tasks
// @Accept json
// @Produce json
func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	controllers.DeleteTask(id)
	c.JSON(id)
	return nil
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c *fiber.Ctx) error {
	res := map[string]interface{}{
	   "data": "Server is up and running",
	}
 
	if err := c.JSON(res); err != nil {
	   return err
	}
 
	return nil
 }
