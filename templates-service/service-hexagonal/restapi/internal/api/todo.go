package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"example.com/servicehex/domain/core"
	"example.com/servicehex/domain/entity"
	"example.com/servicehex/restapi/internal/utils"
	"example.com/servicehex/restapi/spec"
)

func (a *restApi) CreateTodo(c *fiber.Ctx) error {
	createTodo := new(spec.DtoCreateTodo)
	if err := c.BodyParser(createTodo); err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error":  "invalid JSON body",
			"reason": err.Error(),
		})
	}

	userInfo, err := utils.ExtractAndDecodeJwtFiber(c)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{
			"error": "failed to validate userID for authentication",
		})
	}

	todo := entity.Todo{
		ID:        core.TodoID(),
		UserID:    userInfo.UserID,
		Text:      createTodo.Text,
		Deadline:  createTodo.Deadline,
		CreatedAt: c.Context().ConnTime(),
	}

	if err := a.serviceTodo.CreateTodo(c.Context(), todo); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{
			"error": "failed to create todo",
		})
	}

	return c.Status(http.StatusCreated).JSON(todo)
}
