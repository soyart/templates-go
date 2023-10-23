package utils

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"

	"example.com/servicehex/restapi/spec"
)

func NewJwtToken(
	iss string, // userID
	secret []byte,
) (
	string, // Token
	time.Time, // Expiration
	error,
) {
	// TODO: investigate if Local() is actually needed
	exp := time.Now().Add(24 * time.Hour).Local()
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": iss,
		"exp": exp.Unix(),
	})

	// Generate JWT token from claims
	token, err := claims.SignedString(secret)
	if err != nil {
		return token, exp, errors.Wrapf(err, "failed to validate with key %s", secret)
	}

	return token, exp, nil
}

func ExtractAndDecodeJwtFiber(c *fiber.Ctx) (spec.UserLoginInfo, error) {
	user, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return spec.UserLoginInfo{}, errors.New("interface{} conversion to github.com/golang-jwt/jwt/v4.Token failed")
	}

	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return spec.UserLoginInfo{}, errors.New("interface{} conversion to github.com/golang-jwt/jwt/v4.MapClaims failed")
	}

	iss := claims["iss"]
	exp := claims["exp"]

	userUuid, ok := iss.(string)
	if !ok {
		return spec.UserLoginInfo{}, fmt.Errorf("ExtractAndDecodeJwtFiber: interface{} conversion failed for: %v", iss)
	}

	expUnixFloat, ok := exp.(float64)
	if !ok {
		return spec.UserLoginInfo{}, fmt.Errorf("jwt expiration extract: interface{} converstion to float64 failed for: %v %T", exp, exp)
	}

	return spec.UserLoginInfo{
		UserID:     userUuid,
		Expiration: time.Unix(int64(expUnixFloat), 0),
	}, nil
}
