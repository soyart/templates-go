package api

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"example.com/servicehex/domain/core"
	"example.com/servicehex/domain/entity"
	"example.com/servicehex/restapi/internal/utils"
	"example.com/servicehex/restapi/spec"
)

func (a *restApi) Register(c *fiber.Ctx) error {
	registerDto := new(spec.DtoRegisterLoginUser)
	if err := c.BodyParser(registerDto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error":  "invalid JSON body",
			"reason": err.Error(),
		})
	}

	if len(registerDto.Username) == 0 {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": "blank username",
		})
	}

	if len(registerDto.Password) == 0 {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": "blank password",
		})
	}

	if err := core.ValidatePassword([]byte(registerDto.Password)); err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error":  "new password validation failed",
			"reason": err.Error(),
		})
	}

	user := entity.User{
		ID:       core.UserID(registerDto.Username),
		Username: registerDto.Username,
		Password: []byte(registerDto.Password),
	}

	if err := a.serviceUser.Register(c.Context(), user); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{
			"error": fmt.Sprintf("failed to register user %s", user.Username),
		})
	}

	return c.Status(http.StatusCreated).JSON(map[string]interface{}{
		"status": "success",
		"user": map[string]string{
			"username": user.Username,
			"user_id":  user.ID,
		},
	})
}

func (a *restApi) Login(c *fiber.Ctx) error {
	loginDto := new(spec.DtoRegisterLoginUser)
	if err := c.BodyParser(loginDto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error":  "invalid JSON body",
			"reason": err.Error(),
		})
	}

	if len(loginDto.Username) == 0 {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": "blank username",
		})
	}

	user, err := a.serviceUser.Login(c.Context(), loginDto.Username, []byte(loginDto.Password))
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(map[string]string{
			"error": fmt.Sprintf("failed to authenticate user %s", loginDto.Username),
		})
	}

	token, exp, err := utils.NewJwtToken(user.ID, a.authSecret)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{
			"error": "failed to generate authentication token",
		})
	}

	return c.Status(http.StatusOK).JSON(map[string]interface{}{
		"user_id":   user.ID,
		"username":  user.Username,
		"token":     token,
		"token_exp": exp,
	})
}
