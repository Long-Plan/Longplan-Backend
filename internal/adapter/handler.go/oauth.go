package handler

import (
	"log"
	"longplan-backend-service/config"
	"longplan-backend-service/pkg/errors"
	"longplan-backend-service/pkg/lodash"
	"longplan-backend-service/pkg/oauth"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/samber/lo"
)

type oauthHandler struct {
}

func NewOauthHandler() *oauthHandler {
	return &oauthHandler{}
}

func (h oauthHandler) SignIn(c *fiber.Ctx) error {
	config := config.Config.Application
	code := c.Query("code", "")
	if lo.IsEmpty(code) {
		return lodash.ResponseBadRequest(c)
	}
	user, err := oauth.CmuOauthValidation(code)
	if err != nil {
		return lodash.ResponseError(c, errors.NewStatusBadGatewayError(err.Error()))
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, &oauth.UserClaims{
		User: *user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * 12 * time.Hour)),
		},
	})

	token, err := claims.SignedString([]byte(config.Secret))
	if err != nil {
		log.Print(err)
		return lodash.ResponseError(c, errors.NewInternalError(err.Error()))
	}

	// cookie := new(fiber.Cookie)
	// cookie.Name = "token"
	// cookie.Value = token
	// cookie.MaxAge = 1 * 12 * 60 * 60
	// cookie.Path = "/"
	// cookie.HTTPOnly = true
	// cookie.SameSite = "lax"
	// cookie.Domain = config.Domain
	// c.Cookie(cookie)

	return lodash.ResponseOK(c, token)
}

func (h oauthHandler) GetUser(c *fiber.Ctx) error {
	user := c.Locals("USER_DATA").(oauth.UserDto)
	// log.Print(user)
	return lodash.ResponseOK(c, user)
}

func (h oauthHandler) Logout(c *fiber.Ctx) error {
	// c.ClearCookie("token")
	return lodash.ResponseNoContent(c, nil)
}
