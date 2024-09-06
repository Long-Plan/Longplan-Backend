package api

import (
	"longplan-backend-service/internal/adapter/handler.go"
	middlewares "longplan-backend-service/internal/adapter/middleware"

	"github.com/gofiber/fiber/v2"
)

const OAUTH_PREFIX = "/oauth"

func bindOauthRouter(router fiber.Router) {
	oauth := router.Group(OAUTH_PREFIX)

	hdl := handler.NewOauthHandler()
	oauth.Get("/me", middlewares.AuthMiddleware(), hdl.GetUser)

	oauth.Post("", hdl.SignIn)
	oauth.Post("/signout", middlewares.AuthMiddleware(), hdl.Logout)
}
