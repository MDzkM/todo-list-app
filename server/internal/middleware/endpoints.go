package middleware

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func GetAllTask(c *fiber.Ctx) {
	payload := getAllTask()
	c.JSON(payload)
}

func CreateTask(c *fiber.Ctx) {
	var task models.Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	insertOneTask(task)
	c.JSON(task)
}

func TaskComplete(c *fiber.Ctx) {
	id := c.Params("id")
	taskComplete(id)
	c.JSON(id)
}

func DeleteTask(c *fiber.Ctx) {
	id := c.Params("id")
	deleteOneTask(id)
	c.JSON(id)

}
