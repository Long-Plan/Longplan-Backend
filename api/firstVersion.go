package api

import (
	"github.com/gofiber/fiber/v2"
)

func bindFirstVersionRouter(router fiber.Router) {
	firstVersion := router.Group("/v1")
	bindOauthRouter(firstVersion)

	firstVersion.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hi! welcome to LONGPLAN-API 🌈")
	})
	
	bindStudentRouter(firstVersion)
	bindAdminRouter(firstVersion)
	bindCurriculumRouter(firstVersion)
}
