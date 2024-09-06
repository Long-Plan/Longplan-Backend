package api

import "github.com/gofiber/fiber/v2"

const ADMIN_PREFIX = "/admin"

func bindAdminRouter(router fiber.Router) {
	admin := router.Group(ADMIN_PREFIX)
	admin.Get("/home", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Welcome to Admin page!")
	})
	// admin.Get("/getAdmin", handler.GetAdmin)
	// admin.Post("/createAdmin", handler.CreateAdmin)
	// admin.Put("/updateAdmin", handler.UpdateAdmin)
	// admin.Delete("/deleteAdmin", handler.DeleteAdmin)
}