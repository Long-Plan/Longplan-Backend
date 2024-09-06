package middlewares

import (
	"log"
	"longplan-backend-service/config"
	"longplan-backend-service/pkg/errors"
	"longplan-backend-service/pkg/lodash"
	"longplan-backend-service/pkg/oauth"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/samber/lo"
)

func AuthMiddleware() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		config := config.Config.Application

		invalidToken := errors.NewUnauthorizedError(errors.AuthErr("invalid token").Error())

		token := c.Get("Authorization")
		if lo.IsEmpty(token) {
			return lodash.ResponseError(c, errors.NewUnauthorizedError("empty token"))
		}

		// log.Print(token)

		// token := c.Cookies("token")
		// if lo.IsEmpty(token) {
		// 	return lodash.ResponseError(c, errors.NewUnauthorizedError("empty token"))
		// }

		parsedAccessToken, err := jwt.ParseWithClaims(token, &oauth.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Secret), nil
		})
		if err != nil {
			log.Print(err)
			return lodash.ResponseError(c, invalidToken)
		}
		user := &parsedAccessToken.Claims.(*oauth.UserClaims).User

		c.Locals("USER_DATA", *user)
		return c.Next()
	}
}
