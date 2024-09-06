package api

import (
	"longplan-backend-service/internal/adapter/handler.go"

	"github.com/gofiber/fiber/v2"
)

const CURRICULUM_PREFIX = "/curriculum"

func bindCurriculumRouter(router fiber.Router) {
	curriculum := router.Group(CURRICULUM_PREFIX)
	
	curriculum.Get("/", handler.GetAllCurriculum)
	curriculum.Get("byID/:id", handler.GetCurriculumByID)
	curriculum.Get("/old", handler.GetOldCurriculum)
	curriculum.Get("/title", handler.GetMatchedFilenamesHandler)
	curriculum.Get("/course/:id", handler.GetCourseTitle)
	
	curriculum.Post("/parse", handler.ParseCurriculum)
}