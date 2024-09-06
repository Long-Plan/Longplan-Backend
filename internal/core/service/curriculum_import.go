package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// type Course struct {
// 	CourseNo      	string  `json:"courseNo"`
// 	CourseTitleEng 	string	`json:"courseTitleEng"`
// 	RecommendSemester int	`json:"recommendSemester"`
// 	RecommendYear    int	`json:"recommendYear"`
// 	Prerequisites    []string	`json:"prerequisites"`
// 	Corequisite      string	`json:"corequisite"`
// 	Credits          int	`json:"credits"`
// }

// type Group struct {
// 	GroupName      string       `json:"groupName"`
// 	RequiredCourses []Course    `json:"requiredCourses"`
// 	ElectiveCourses []Course    `json:"electiveCourses"`
// }

// type CurriculumData struct {
// 	RequiredCredits     int     `json:"requiredCredits"`
// 	FreeElectiveCredits int     `json:"freeElectiveCredits"`
// 	CoreAndMajorGroups  []Group `json:"coreAndMajorGroups"`
// 	GeGroups            []Group `json:"geGroups"`
// }

type New_Curriculum struct {
	CoreGroup 	[]New_Group `json:"coreGroup"`
	MajorGroup 	[]New_Group `json:"majorGroup"`
	GeGroup 	[]New_Group `json:"geGroup"`
}

type New_Group struct {
	GroupName      	string       	`json:"groupName"`
	RequiredCourses []New_Course    `json:"requiredCourses"`
	ElectiveCourses []New_Course    `json:"electiveCourses"`
}

type New_Course struct {
	CourseNo      		string  	`json:"courseNo"`
	RecommendSemester 	int			`json:"recommendSemester"`
	RecommendYear    	int			`json:"recommendYear"`
	Prerequisites    	[]string	`json:"prerequisites"`
	Corequisite      	string		`json:"corequisite"`
}

// Convert an old Course to a New_Course
func mapCourse(oldCourse Course) New_Course {
	return New_Course{
		CourseNo: oldCourse.CourseNo,
		RecommendSemester: oldCourse.RecommendSemester,
		RecommendYear: oldCourse.RecommendYear,
		Prerequisites: oldCourse.Prerequisites,
		Corequisite: oldCourse.Corequisite,
	}
}

// Convert an old Group to a New_Group
func mapGroup(oldGroup Group) New_Group {
	newRequiredCourses := make([]New_Course, len(oldGroup.RequiredCourses))
	for i, course := range oldGroup.RequiredCourses {
		newRequiredCourses[i] = mapCourse(course)
	}

	newElectiveCourses := make([]New_Course, len(oldGroup.ElectiveCourses))
	for i, course := range oldGroup.ElectiveCourses {
		newElectiveCourses[i] = mapCourse(course)
	}

	return New_Group{
		GroupName: oldGroup.GroupName,
		RequiredCourses: newRequiredCourses,
		ElectiveCourses: newElectiveCourses,
	}
}

// Convert old CurriculumData to New_Curriculum
func mapCurriculum(oldCurriculum CurriculumData) New_Curriculum {
	coreGroup := []New_Group{}
	majorGroup := []New_Group{}

	for _, group := range oldCurriculum.CoreAndMajorGroups {
		if group.GroupName == "Core" {
			coreGroup = append(coreGroup, mapGroup(group))
		} else {
			majorGroup = append(majorGroup, mapGroup(group))
		}
	}

	geGroup := make([]New_Group, len(oldCurriculum.GeGroups))
	for i, group := range oldCurriculum.GeGroups {
		geGroup[i] = mapGroup(group)
	}

	return New_Curriculum{
		CoreGroup: coreGroup,
		MajorGroup: majorGroup,
		GeGroup: geGroup,
	}
}

func ParseJSON() {
	// Load JSON data from file
	file, err := os.Open("./internal/core/service/data/CPE-2563-normal.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	var oldCurriculum CurriculumData

	// Parse JSON data into old curriculum structure
	if err := json.Unmarshal(byteValue, &oldCurriculum); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Map old curriculum to new curriculum structure
	newCurriculum := mapCurriculum(oldCurriculum)

	// Convert the new curriculum structure to JSON
	newCurriculumJSON, err := json.MarshalIndent(newCurriculum, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling new curriculum to JSON:", err)
		return
	}

	// Write the new JSON to a file
	err = ioutil.WriteFile("./internal/core/service/data/New-CPE-2563-structure.json", newCurriculumJSON, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}

	fmt.Println("New curriculum JSON file created successfully: new_curriculum.json")
}
