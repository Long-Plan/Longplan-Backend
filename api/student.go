package api

import (
	"longplan-backend-service/internal/adapter/handler.go"
	middlewares "longplan-backend-service/internal/adapter/middleware"

	"github.com/gofiber/fiber/v2"
)

const STUDENT_PREFIX = "/student"

func bindStudentRouter(router fiber.Router) {
	enrolledData := router.Group(STUDENT_PREFIX)

	// main method
	enrolledData.Get("/enrolledData",middlewares.AuthMiddleware(), handler.ScrapingHandler)
	enrolledData.Get("/enrolledDataOld",middlewares.AuthMiddleware(), handler.ScrapingOldHandler)
	
	// adv-project
	// enrolledData.Get("/enrolledData/get/cpestudents", handler.GetCPEstudents) 	
	// enrolledData.Get("/enrolledData/:id", handler.GetEnrolledDataHandler) 		
	// enrolledData.Get("/enrolledData/course/:id", handler.GetCourseTitle)
}
