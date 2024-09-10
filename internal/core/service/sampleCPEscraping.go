package service

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type EnrollCourse struct {
	Course_no   string
	CourseTitle string
	Credit      string
	Grade       string
}

type Students struct {
	StudentID string
	EnrollCourse []MappingEnrolledCourseAdv
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	return i
}

func MultipleScrapingAndsaveToJSON(start string, end string, normal bool) {
	// start 619 - end 676
	allEnrollCourses := make(map[string][]MappingEnrolledCourseAdv)

	intStart := toInt(start)
	intEnd := toInt(end)


	if normal {
		for i := intStart; i <= intEnd; i++ {
			studentID := "640610" + fmt.Sprintf("%03d", i)
			studentEnrollCourse , err := Scraping_Enrollcourse_adv(studentID)
			if err != nil {
				fmt.Println("Error Scraping_Enrollcourse:", err)
				return
			}
			allEnrollCourses[studentID] = studentEnrollCourse
		}
	} else {
		for i := intStart; i <= intEnd; i++ {
			studentID := "640612" + fmt.Sprintf("%03d", i)
			studentEnrollCourse , err := Scraping_Enrollcourse_adv(studentID)
			if err != nil {
				fmt.Println("Error Scraping_Enrollcourse:", err)
				return
			}
			allEnrollCourses[studentID] = studentEnrollCourse
		}
	}

	var special string
	if normal {
		special = "normal"
	} else {
		special = "special"
	}
	path := fmt.Sprintf("internal/core/service/data/students_enrolled_course_%s_%s-%s.json", special, start, end)
	// Save the allEnrollCourses slice to a JSON file
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error creating JSON file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Set indentation for pretty printing

	err = encoder.Encode(allEnrollCourses)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	fmt.Println("JSON file created successfully")
}