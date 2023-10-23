package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"

	"example.com/servicehex/adapter/datagateway/redisadapter"
	"example.com/servicehex/adapter/service/todo"
	"example.com/servicehex/adapter/service/user"
	"example.com/servicehex/config"
	"example.com/servicehex/domain/datagateway/mock"
	"example.com/servicehex/restapi/internal/api"
	"example.com/servicehex/restapi/internal/middleware"
)

func main() {
	confPath, found := os.LookupEnv("CONF_PATH")
	if !found {
		confPath = "./config"
	}

	conf, err := config.LoadConfig(confPath)
	if err != nil {
		log.Fatalf("failed to init config: %s", err.Error())
	}

	// redisadpter does not implement
	// datagateway.DataGatewayUser **yet**.
	repoRedis := redisadapter.New(conf.Redis)
	if repoRedis == nil {
		log.Fatalf("nil redis repository")
	}

	mockRepoUser := mock.NewMockRepoUser()

	serviceTodo := todo.New(repoRedis)
	serviceUser := user.New(mockRepoUser)

	api := api.New(serviceTodo, serviceUser)

	app := fiber.New()
	jwtMiddleware := middleware.JwtAuthentication(conf.Rest.JwtSecret)

	groupTodo := app.Group("/todos")
	groupTodo.Use(jwtMiddleware)
	groupTodo.Post("/", api.CreateTodo)

	groupUser := app.Group("/users")
	groupUser.Post("/login", api.Login)
	groupUser.Post("/register", api.Register)

	// @TODO: Graceful shutdown
	app.Listen(conf.Rest.Address)
}
