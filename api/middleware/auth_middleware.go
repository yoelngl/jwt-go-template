package middleware

import (
	"chatapp/model"
	"chatapp/pkg/constant"
	"chatapp/pkg/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware(c *fiber.Ctx) error {
	authorization := c.Get("Authorization")
	if authorization == "" {
		return c.Status(constant.OK_CODE).JSON(fiber.Map{
			"status":  constant.BAD_REQUEST_CODE,
			"message": constant.UNAUTHORIZED,
			"data":    nil,
		})
	}

	jwtString := strings.Split(authorization, "Bearer ")[1]

	claims := &model.JWT{}
	token, err := jwt.ParseWithClaims(jwtString, claims, func(t *jwt.Token) (interface{}, error) {
		return utils.JwtSecretKey, nil
	})

	if err != nil {
		return c.Status(constant.OK_CODE).JSON(fiber.Map{
			"status":  constant.UNAUTHORIZED_CODE,
			"message": constant.INVALID_JWT,
			"data":    nil,
		})
	}

	if !token.Valid {
		return c.Status(constant.OK_CODE).JSON(fiber.Map{
			"status":  constant.UNAUTHORIZED_CODE,
			"message": constant.INVALID_JWT,
			"data":    nil,
		})
	}

	c.Locals("username", claims.Username)
	return c.Next()
}
