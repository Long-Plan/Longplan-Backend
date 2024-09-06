package handler

import (
	"log"
	"longplan-backend-service/internal/core/service"
	"longplan-backend-service/pkg/oauth"
	"regexp"

	"github.com/gofiber/fiber/v2"
)

func ScrapingHandler(c *fiber.Ctx) error { // Next time this method with fetch student_id from the Authorization header
	studentID := c.Locals("USER_DATA").(oauth.UserDto).StudentID
	studentIDRegex := regexp.MustCompile(`^\d{9}$`)

	if studentID == "" {
		return c.Status(fiber.StatusBadRequest).SendString("student_id parameter is required")
	}
	if !studentIDRegex.MatchString(studentID) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid student_id format")
	}

	log.Print(studentID)

	mappings := service.GetEnrolledCourseByStudentID(studentID)
	
	return c.JSON(mappings)
}

func GetEnrolledDataHandler(c *fiber.Ctx) error {
	studentID := c.Params("id")
	studentIDRegex := regexp.MustCompile(`^\d{9}$`)

	if studentID == "" {
		return c.Status(fiber.StatusBadRequest).SendString("student_id parameter is required")
	}
	if !studentIDRegex.MatchString(studentID) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid student_id format")
	}

	log.Print(studentID)

	mappings, err := service.Scraping_Enrollcourse_adv(studentID , true)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(mappings)
}

func GetCourseTitle (c *fiber.Ctx) error {
	courseID := c.Params("id")
	courseIDRegex := regexp.MustCompile(`^\d{6}$`)

	if courseID == "" {
		return c.Status(fiber.StatusBadRequest).SendString("course_id parameter is required")
	}
	if !courseIDRegex.MatchString(courseID) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid course_id format")
	}

	log.Print(courseID)

	courseTitle, err := service.Mapping_coursename_byCourseID(courseID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(courseTitle)
} 

func GetCPEstudents (c *fiber.Ctx) error {
	
	start := "619"
	end := "676"
	normal := true

	service.MultipleScrapingAndsaveToJSON(start, end, normal)

	return c.SendString("Scraping and save to JSON file successfully")

}
