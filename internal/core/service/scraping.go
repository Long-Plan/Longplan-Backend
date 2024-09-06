package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"longplan-backend-service/logs"
	"longplan-backend-service/pkg/errors"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Course struct {
	CourseNo      	string  `json:"courseNo"`	
	CourseTitleEng 	string	`json:"courseTitleEng"`
	RecommendSemester int	`json:"recommendSemester"`
	RecommendYear    int	`json:"recommendYear"`
	Prerequisites    []string	`json:"prerequisites"`
	Corequisite      string	`json:"corequisite"`
	Credits          int	`json:"credits"`
}

type Group struct {
	GroupName      string       `json:"groupName"`
	RequiredCourses []Course    `json:"requiredCourses"`
	ElectiveCourses []Course    `json:"electiveCourses"`
}

type CurriculumData struct {
	RequiredCredits     int     `json:"requiredCredits"`
	FreeElectiveCredits int     `json:"freeElectiveCredits"`
	CoreAndMajorGroups  []Group `json:"coreAndMajorGroups"`
	GeGroups            []Group `json:"geGroups"`
}

type Enrolled_course struct {
	Course_no   string
	Credit      string
	Grade       string
}

type Mapping_Enrolled_course struct {
	Year     string
	Semester string
	Courses  []Enrolled_course
}

type Enrolled_course_adv struct {
	Course_no   string
	CourseTitle string 
	Credit      string
	Grade       string
}

type Mapping_Enrolled_course_adv struct {
	Year     string
	Semester string
	Courses  []Enrolled_course_adv
}

type CourseData struct {
	CourseNo       string `json:"courseNo"`
	CourseTitleEng string `json:"CourseTitleEng"`
	Abbreviation   string `json:"Abbreviation"`
}

func TransformInput(input string) string {
	// Define a regular expression pattern to extract the first two digits.
	re := regexp.MustCompile(`^(\d{2})`)

	// Find the first two digits in the input.
	matches := re.FindStringSubmatch(input)

	if len(matches) != 2 {
		// If the pattern is not found, return an error or handle the case accordingly.
		return "Invalid input"
	}

	// Extract the first two digits and concatenate with "25".
	firstTwoDigits := matches[1]
	result := "25" + firstTwoDigits

	// Convert the result to an integer and then back to a string to remove leading zeros.
	resultInt, _ := strconv.Atoi(result)
	result = strconv.Itoa(resultInt)

	return result
}

// https://reg.eng.cmu.ac.th/reg/plan_detail/plan_data_term.php?student_id=640612093

func read_http(student_id string) (string, error) {
	url := fmt.Sprintf("https://reg.eng.cmu.ac.th/reg/plan_detail/plan_data_term.php?student_id=%v", student_id)
	resp, err := http.Get(url)
	if err != nil {
		logs.Error(err)
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error(err)
		return "", err
	}

	return string(body), nil
}

func Scraping_Enrollcourse(studentID string) ([]Mapping_Enrolled_course, error) {
	body, err := read_http(studentID)
	if err != nil {
		logs.Error(err)
		return nil, errors.NewNotFoundError("Student ID not found")
	}

	var year, semester string
	var mappings []Mapping_Enrolled_course

	query, _ := goquery.NewDocumentFromReader(strings.NewReader(body))
	query.Find("table[width='100%'][border='0'][class='t']").Each(func(i int, s *goquery.Selection) {
		s.Find("td[align='center'] > B").Each(func(j int, row *goquery.Selection) {
			text := row.Text()
			if strings.Contains(text, "ภาคเรียนที่") && strings.Contains(text, "ปีการศึกษา") {
				semesterMatches := regexp.MustCompile(`\d+`).FindAllString(text, -1)
				if len(semesterMatches) >= 2 {
					semester = semesterMatches[len(semesterMatches)-2]
					tempYear := semesterMatches[len(semesterMatches)-1]

					numYear, err := strconv.Atoi(tempYear)
					if err != nil {
						logs.Error(err)
						return
					}

					numStudentID, err := strconv.Atoi(TransformInput(studentID))
					if err != nil {
						logs.Error(err)
						return
					}

					numYear = numYear - numStudentID + 1
					year = strconv.Itoa(numYear)
				}
			}
		})

	var currentCourses []Enrolled_course
	
		s.Find("table[cellspacing='1'][cellpadding='3'][width='60%'][border='0'][class='t'] tr[bgcolor='#FFFFFF']").Each(func(j int, row *goquery.Selection) {
			course_no := strings.TrimSpace(row.Find("td:first-child").Text())
			credit := strings.TrimSpace(row.Find("td:nth-child(2)").Text())
			grade := strings.TrimSpace(row.Find("td:nth-child(3)").Text())
			
				currentCourses = append(currentCourses, Enrolled_course{
					Course_no:   course_no,
					Credit:      credit,
					Grade:       grade,
				})
		})

		if len(currentCourses) > 0 {
				mappings = append(mappings, Mapping_Enrolled_course{
					Year:     year,
					Semester: semester,
					Courses:  currentCourses,
				})
			}
		
	})

	return mappings, nil
}

func Scraping_Enrollcourse_adv(studentID string , Mode bool) ([]Mapping_Enrolled_course_adv, error) {
	body, err := read_http(studentID)
	if err != nil {
		logs.Error(err)
		return nil, errors.NewNotFoundError("Student ID not found")
	}

	var year, semester string
	var mappings []Mapping_Enrolled_course_adv

	query, _ := goquery.NewDocumentFromReader(strings.NewReader(body))
	query.Find("table[width='100%'][border='0'][class='t']").Each(func(i int, s *goquery.Selection) {
		s.Find("td[align='center'] > B").Each(func(j int, row *goquery.Selection) {
			text := row.Text()
			if strings.Contains(text, "ภาคเรียนที่") && strings.Contains(text, "ปีการศึกษา") {
				semesterMatches := regexp.MustCompile(`\d+`).FindAllString(text, -1)
				if len(semesterMatches) >= 2 {
					semester = semesterMatches[len(semesterMatches)-2]
					tempYear := semesterMatches[len(semesterMatches)-1]

					numYear, err := strconv.Atoi(tempYear)
					if err != nil {
						logs.Error(err)
						return
					}

					numStudentID, err := strconv.Atoi(TransformInput(studentID))
					if err != nil {
						logs.Error(err)
						return
					}

					numYear = numYear - numStudentID + 1
					year = strconv.Itoa(numYear)
				}
			}
		})

	var currentCourses []Enrolled_course_adv
	
		s.Find("table[cellspacing='1'][cellpadding='3'][width='60%'][border='0'][class='t'] tr[bgcolor='#FFFFFF']").Each(func(j int, row *goquery.Selection) {
			course_no := strings.TrimSpace(row.Find("td:first-child").Text())
			credit := strings.TrimSpace(row.Find("td:nth-child(2)").Text())
			grade := strings.TrimSpace(row.Find("td:nth-child(3)").Text())
			
			if Mode{
				courseTitle, err := findCourseTitle(course_no) 
				if err != nil {
					logs.Error(err)
					return
				}
			currentCourses = append(currentCourses, Enrolled_course_adv{
				Course_no:   course_no,
				CourseTitle: courseTitle,
				Credit:      credit,
				Grade:       grade,
			})
			} else {
				currentCourses = append(currentCourses, Enrolled_course_adv{
					Course_no:   course_no,
					Credit:      credit,
					Grade:       grade,
				})
			}
		})

		if len(currentCourses) > 0 {
				mappings = append(mappings, Mapping_Enrolled_course_adv{
					Year:     year,
					Semester: semester,
					Courses:  currentCourses,
				})
			}
		
	})

	return mappings, nil
}

func Mapping_coursename_byCourseID(courseID string) (string, error) {
	url := fmt.Sprintf("https://mis-api.cmu.ac.th/tqf/v1/course-template?courseid=%v&academicyear=2563&academicterm=1", courseID)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var courses []CourseData
	err = json.Unmarshal(body, &courses)
	if err != nil {
		return "", err
	}

	var courseTitles string

	if len(courses) > 0 {
		if courses[0].Abbreviation != "" {
			courseTitles = courses[0].Abbreviation // if the course has an abbreviation
		} else {
			courseTitles = courses[0].CourseTitleEng // if the course doesn't have an abbreviation, use the long name
		}
	}

	return courseTitles, nil
}



func findCourseTitle(courseNo string) (courseTitleEng string, err error) {
	curriculumData, err := getCurriculum(CurriculumPayload{
		Major: "CPE",
		Year:  "2563",
		Plan:  "normal",
	})
	if err != nil {
		logs.Error(err)
		return "", err
	} 
	if curriculumData != nil {
		// Iterate through the required courses in geGroups and coreAndMajorGroups
		for _, group := range append(curriculumData.GeGroups, curriculumData.CoreAndMajorGroups...) {
			for _, course := range group.RequiredCourses {
				if course.CourseNo == courseNo {
					return course.CourseTitleEng, nil
				} else {
					for _, electiveCourse := range group.ElectiveCourses {
						if electiveCourse.CourseNo == courseNo {
							return electiveCourse.CourseTitleEng, nil
						}
					}
				}
			}
		}
	}
	return "Not contain in curriculum", nil
}

type CurriculumPayload struct {
	Major string
	Year  string
	Plan  string
}


func getCurriculum(payload CurriculumPayload) (*CurriculumData, error) {
	path := fmt.Sprintf("internal/core/service/data/%s-%s-%s.json", payload.Major, payload.Year, payload.Plan)

	// Open the file
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Create a timeout context if needed
	timeout := time.After(10 * time.Second)
	done := make(chan error, 1)

	var curriculumData CurriculumData

	go func() {
		// Decode the JSON file into the struct
		done <- json.NewDecoder(file).Decode(&curriculumData)
	}()

	select {
	case <-timeout:
		return nil, fmt.Errorf("timed out while reading the file")
	case err := <-done:
		if err != nil {
			return nil, fmt.Errorf("failed to decode JSON: %w", err)
		}
	}

	return &curriculumData, nil
}


type enrolledCourse struct {
	Year         string `json:"year"`
	Semester     string `json:"semester"`
	CourseNumber string `json:"courseNo"`
	Credit       string `json:"credit"`
	Grade        string `json:"grade"`
}

func getEnrolledCourses(studentID string) ([]enrolledCourse, error) {
	body, err := read_http(studentID)
	if err != nil {
		logs.Error(err)
		return nil, errors.NewNotFoundError("Student ID not found")
	}

	// Parse the document from the response body.
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed to parse document: %v", err)
	}
	var courses []enrolledCourse
	var year, semester string

	// Find and process each block representing courses for a year and semester
	doc.Find("table[width='100%'][border='0'][class='t']").Each(func(i int, s *goquery.Selection) {
		// Process each row in the table
		s.Find("td[align='center'] > B").Each(func(j int, row *goquery.Selection) {
			// Check if the text contains Thai characters for semester and year
			if strings.Contains(row.Text(), "ภาคเรียนที่") && strings.Contains(row.Text(), "ปีการศึกษา") {
				// Extract semester and year from the text
				semesterMatches := regexp.MustCompile(`\d+`).FindAllString(row.Text(), -1)
				if len(semesterMatches) >= 2 {
					semester = semesterMatches[len(semesterMatches)-2]
					temp_year := semesterMatches[len(semesterMatches)-1]

					num_year, nil := strconv.Atoi(temp_year)
					if err != nil {
						fmt.Println("Error converting year to int:", err)
						return
					}
					num_studentID, nil := strconv.Atoi(transformInput(studentID))
					if err != nil {
						fmt.Println("Error converting student to int:", err)
						return
					}

					num_year = num_year - num_studentID + 1 // calculated study year by studentID and academic year
					year = strconv.Itoa(num_year)

				}
			}

		})
		s.Find("table[cellspacing='1'][cellpadding='3'][width='60%'][border='0'][class='t'] tr[bgcolor='#FFFFFF']").Each(func(j int, row *goquery.Selection) {
			// Extract course details from each row
			courseNumber := strings.TrimSpace(row.Find("td:first-child").Text())
			credit := strings.TrimSpace(row.Find("td:nth-child(2)").Text())
			grade := strings.TrimSpace(row.Find("td:nth-child(3)").Text())

			// Add the course details to the courses slice
			courses = append(courses, enrolledCourse{
				CourseNumber: courseNumber,
				Credit:       credit,
				Grade:        grade,
				Semester:     semester,
				Year:         year,
			})
		})

	})

	return courses, nil
}

func groupCoursesByYearSemester(courses []enrolledCourse) map[string]map[string][]enrolledCourse {
	groupedCourses := make(map[string]map[string][]enrolledCourse)

	for _, course := range courses {
		if _, ok := groupedCourses[course.Year]; !ok {
			groupedCourses[course.Year] = make(map[string][]enrolledCourse)
		}
		groupedCourses[course.Year][course.Semester] = append(groupedCourses[course.Year][course.Semester], course)
	}

	// Ensure at least 4 years of data
	for year := 1; year <= 4; year++ {
		yearStr := strconv.Itoa(year)
		if _, ok := groupedCourses[yearStr]; !ok {
			groupedCourses[yearStr] = make(map[string][]enrolledCourse)
		}
		// Ensure there are two semesters for each year
		if _, ok := groupedCourses[yearStr]["1"]; !ok {
			groupedCourses[yearStr]["1"] = []enrolledCourse{}
		}
		if _, ok := groupedCourses[yearStr]["2"]; !ok {
			groupedCourses[yearStr]["2"] = []enrolledCourse{}
		}
	}

	return groupedCourses
}

func transformInput(input string) string {
	// Define a regular expression pattern to extract the first two digits.
	re := regexp.MustCompile(`^(\d{2})`)

	// Find the first two digits in the input.
	matches := re.FindStringSubmatch(input)

	if len(matches) != 2 {
		// If the pattern is not found, return an error or handle the case accordingly.
		return "Invalid input"
	}

	// Extract the first two digits and concatenate with "25".
	firstTwoDigits := matches[1]
	result := "25" + firstTwoDigits

	// Convert the result to an integer and then back to a string to remove leading zeros.
	resultInt, _ := strconv.Atoi(result)
	result = strconv.Itoa(resultInt)

	return result
}

func GetEnrolledCourseByStudentID(id string) map[string]map[string][]enrolledCourse {
	// Input your student ID here
	studentID := id

	courses, err := getEnrolledCourses(studentID)
	if err != nil {
		log.Fatal(err)
	}
	groupedCourses := groupCoursesByYearSemester(courses)

	return groupedCourses
	// writeToFile(courses,studentID)
}