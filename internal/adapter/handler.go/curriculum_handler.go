package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"longplan-backend-service/internal/core/service"
	"os"
	"path/filepath"
	"regexp"

	"github.com/gofiber/fiber/v2"
)

// GetCurriculum retrieves curriculum data for all matched majors and years from JSON files
func GetAllCurriculum(c *fiber.Ctx) error {
    // Retrieve all curriculum data from matched filenames
    responseStr, err := GetMatchedFilenames("internal/core/service/data/")
    if err != nil {
        // Return an error response if the function fails
        return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving matched filenames: " + err.Error())
    }

    var response Response
    err = json.Unmarshal([]byte(responseStr), &response)
    if err != nil {
        // Return an error response if unable to unmarshal JSON
        return c.Status(fiber.StatusInternalServerError).SendString("Error unmarshalling JSON: " + err.Error())
    }

    // Check if there are any curriculums
    if len(response.Filenames) == 0 {
        return c.Status(fiber.StatusNotFound).SendString("No curriculums found")
    }

    // Slice to hold all curriculum data
    var allCurriculumsData []interface{}

    // Iterate over each curriculum to read its corresponding JSON file
    for _, curr := range response.Filenames {
        filename := "./internal/core/service/data/" + curr.Major + "_" + curr.Year + ".json"

        // Read the JSON file using the ReadJSONFile function
        jsonFile, err := ReadJSONFile(filename)
        if err != nil {
            // Log the error for this specific file and continue
            // Optionally, you can return an error here or skip to the next file
            return c.Status(fiber.StatusInternalServerError).SendString("Error reading JSON file: " + filename)
        }

        // Append the read JSON data to the allCurriculumsData slice
        allCurriculumsData = append(allCurriculumsData, jsonFile)
    }

    // Return all curriculum data as JSON response
    return c.JSON(allCurriculumsData)
}

func GetCurriculumByID(c *fiber.Ctx) error {
	id , err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID format")
	}

	// Retrieve all curriculum data from matched filenames
    responseStr, err := GetMatchedFilenames("internal/core/service/data/")
    if err != nil {
        // Return an error response if the function fails
        return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving matched filenames: " + err.Error())
    }

    var response Response
    err = json.Unmarshal([]byte(responseStr), &response)
    if err != nil {
        // Return an error response if unable to unmarshal JSON
        return c.Status(fiber.StatusInternalServerError).SendString("Error unmarshalling JSON: " + err.Error())
    }

    // Check if there are any curriculums
    if len(response.Filenames) == 0 {
        return c.Status(fiber.StatusNotFound).SendString("No curriculums found")
    }

    // Slice to hold all curriculum data
    var allCurriculumsData []interface{}

    // Iterate over each curriculum to read its corresponding JSON file
    for _, curr := range response.Filenames {
        filename := "./internal/core/service/data/" + curr.Major + "_" + curr.Year + ".json"

        // Read the JSON file using the ReadJSONFile function
        jsonFile, err := ReadJSONFile(filename)
        if err != nil {
            // Log the error for this specific file and continue
            // Optionally, you can return an error here or skip to the next file
            return c.Status(fiber.StatusInternalServerError).SendString("Error reading JSON file: " + filename)
        }

        // Append the read JSON data to the allCurriculumsData slice
        allCurriculumsData = append(allCurriculumsData, jsonFile)
    }

	curriculumCount := len(allCurriculumsData)
	if id >= curriculumCount {
		return c.Status(fiber.StatusNotFound).SendString("Curriculum not found")
	}

	return c.JSON(allCurriculumsData[id])
}

func GetOldCurriculum(c *fiber.Ctx) error {
		major := c.Query("major")
		year := c.Query("year")
		plan := c.Query("plan")
		filename := "./internal/core/service/data/" + major + "-" + year + "-" + plan + ".json"

		// Read JSON file using the function
		jsonFile, err := ReadJSONFile(filename)
		if err != nil {
			// Return an error response if unable to read the file
			return c.Status(fiber.StatusInternalServerError).SendString("Error reading JSON file : " + filename)
		}

		// Return curriculum data as JSON response
		return c.JSON(jsonFile)
}

func GetGEtemplate(c *fiber.Ctx) error {
	major := c.Query("major")
	year := c.Query("year")
	filename := "./internal/core/service/data/" + major + "_GE_" + year + ".json"

	jsonFile, err := ReadJSONFile(filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error reading JSON file : " + filename)
	}

	return c.JSON(jsonFile)
}

func ReadJSONFile(filePath string) (interface{}, error) {
	// Read the JSON file
	jsonFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		// Return error if unable to read the file
		return nil, err
	}

	// Initialize a Curriculum struct to hold the parsed JSON data
	var curriculum interface{}

	// Unmarshal the JSON data into the Curriculum struct
	err = json.Unmarshal(jsonFile, &curriculum)
	if err != nil {
		// Return error if unable to unmarshal JSON
		return nil, err
	}

	// Return the parsed Curriculum struct
	return &curriculum, nil
}

func ParseCurriculum (c *fiber.Ctx) error {
	service.ParseJSON()
	return c.SendString("Curriculum JSON file parsed successfully")
}

type Response struct {
    Filenames []curriculum `json:"curriculum"`
}

type curriculum struct {
	Major 	string `json:"major"`
	Year 	string `json:"year"`
}

// GetMatchedFilenames returns a JSON string with the majors and years extracted from filenames matching the regex pattern in the specified directory
func GetMatchedFilenames(directory string) (string, error) {
    // Slice to store curriculum data
    var curriculums []curriculum

    // Compile the regular expression
    re := regexp.MustCompile(`^([A-Z]+)_([0-9]+)`)

    // Walk the directory
    err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        // Check if it’s a file
        if !info.IsDir() {
            filename := info.Name()

            // Apply regex matching
            if matches := re.FindStringSubmatch(filename); matches != nil {
                // Create a curriculum instance
                curriculums = append(curriculums, curriculum{
                    Major: matches[1], // First capture group (major)
                    Year:  matches[2], // Second capture group (year)
                })
            }
        }
        return nil
    })

    if err != nil {
        return "", fmt.Errorf("error walking the path: %v", err)
    }

    // Prepare the response
    response := Response{
        Filenames: curriculums,
    }

    // Convert response to JSON
    jsonResponse, err := json.MarshalIndent(response, "", "  ")
    if err != nil {
        return "", fmt.Errorf("error marshalling JSON: %v", err)
    }

    // Return JSON as a string
    return string(jsonResponse), nil
}


func GetMatchedFilenamesHandler(c *fiber.Ctx) error {
	// Get the directory path from the query parameter
	directory := "internal/core/service/data/"

	// Call the GetMatchedFilenames function
	jsonResponse, err := GetMatchedFilenames(directory)
	if err != nil {
		// Return an error response if the function fails
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// Return the JSON response
	return c.SendString(jsonResponse)
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