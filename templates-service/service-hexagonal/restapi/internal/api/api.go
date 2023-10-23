package api

import (
	"github.com/gofiber/fiber/v2"

	"example.com/servicehex/domain/core"
)

type RestApi interface {
	CreateTodo(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
}

func New(serviceTodo core.PortTodo, serviceUser core.PortUser) RestApi {
	return &restApi{
		serviceTodo: serviceTodo,
		serviceUser: serviceUser,
	}
}

// A HTTP REST `view/presentation` layer of our program
// Adapter `serviceTodo` connects to core port PortTodo,
// likewise, `serviceUser` connects to PortUser.
type restApi struct {
	serviceTodo core.PortTodo
	serviceUser core.PortUser

	authSecret []byte
}
