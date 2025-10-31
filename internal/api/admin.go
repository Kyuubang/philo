//go:build admin

package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type UsersResponse struct {
	Class    string    `json:"class"`
	Students []Student `json:"students"`
}

type Student struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

type GetClassesResponse struct {
	Classes []Class `json:"classes"`
}

type Class struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (server CustomServer) GetUsers(classId string) (response UsersResponse, code int, err error) {
	// create new request
	req, err := http.NewRequest("GET", server.Host+"/v1/admin/user", nil)
	if err != nil {
		return response, -1, err
	}

	// add query
	q := req.URL.Query()       // Get a copy of the query values.
	q.Add("class_id", classId) // Add a new value to the set.
	req.URL.RawQuery = q.Encode()

	// set header
	req.Header.Set("Authorization", "Bearer "+server.Token)
	req.Header.Set("Content-Type", "application/json")

	// send request
	client := getHTTPClient()
	resp, err := client.Do(req)
	if err != nil {
		return response, -1, err
	}

	defer resp.Body.Close()

	// read response code
	code = resp.StatusCode

	// read response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, code, err
	}
	// filter response output
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return response, code, err
	}

	return response, code, nil
}

// GetClasses is a function to get all class
func (server CustomServer) GetClasses() (response GetClassesResponse, code int, err error) {
	// create new request
	req, err := http.NewRequest("GET", server.Host+"/v1/admin/class", nil)
	if err != nil {
		return response, -1, err
	}

	// set header
	req.Header.Set("Authorization", "Bearer "+server.Token)
	req.Header.Set("Content-Type", "application/json")

	// send request
	client := getHTTPClient()
	resp, err := client.Do(req)
	if err != nil {
		return response, -1, err
	}

	defer resp.Body.Close()

	// read response code
	code = resp.StatusCode

	// read response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, code, err
	}
	// filter response output
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return response, code, err
	}

	return response, code, nil
}

// CreateClass is a function to create new class
func (server CustomServer) CreateClass(className string) (response CreateClassResponse, code int, err error) {
	// create json data
	jsonData, err := json.Marshal(map[string]string{"name": className})
	if err != nil {
		return response, code, err
	}

	// create request
	req, err := http.NewRequest("POST", server.Host+"/v1/admin/class", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+server.Token)

	// send request
	client := getHTTPClient()
	resp, err := client.Do(req)
	if err != nil {
		return response, code, err
	}

	defer resp.Body.Close()

	// read response code
	code = resp.StatusCode

	// read response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, code, err
	}
	// filter response output
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return response, code, err
	}

	return response, code, nil
}

// UpdateClass is a function to update class
func (server CustomServer) UpdateClass(classId int, className string) (response map[string]string, code int, err error) {
	// create json data
	jsonData, err := json.Marshal(map[string]string{"name": className})
	if err != nil {
		return response, code, err
	}

	// create request
	req, err := http.NewRequest("PUT", server.Host+"/v1/admin/class", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+server.Token)

	// convert int to string
	classIdStr := strconv.Itoa(classId)

	// add query
	q := req.URL.Query()          // Get a copy of the query values.
	q.Add("class_id", classIdStr) // Add a new value to the set.
	req.URL.RawQuery = q.Encode()

	// send request
	client := getHTTPClient()
	resp, err := client.Do(req)
	if err != nil {
		return response, code, err
	}

	defer resp.Body.Close()

	// read response code
	code = resp.StatusCode

	// read response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, code, err
	}
	// filter response output
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return response, code, err
	}

	return response, code, nil
}

// DeleteClass is a function to delete class
func (server CustomServer) DeleteClass(classId int) (response map[string]string, code int, err error) {
	// create request
	req, err := http.NewRequest("DELETE", server.Host+"/v1/admin/class", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+server.Token)

	// convert int to string
	classIdStr := strconv.Itoa(classId)

	// add query
	q := req.URL.Query()          // Get a copy of the query values.
	q.Add("class_id", classIdStr) // Add a new value to the set.
	req.URL.RawQuery = q.Encode()

	// send request
	client := getHTTPClient()
	resp, err := client.Do(req)
	if err != nil {
		return response, code, err
	}

	defer resp.Body.Close()

	// read response code
	code = resp.StatusCode

	// read response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, code, err
	}
	// filter response output
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return response, code, err
	}

	return response, code, nil
}

// CreateLabs is a function to create new lab
func (server CustomServer) CreateLabs(courseId int, labName string) (response CreateLabsResponse, code int, err error) {
	// create json data
	jsonData, err := json.Marshal(map[string]interface{}{"name": labName, "course_id": courseId})
	if err != nil {
		return response, code, err
	}

	// create request
	req, err := http.NewRequest("POST", server.Host+"/v1/admin/labs", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+server.Token)

	// send request
	client := getHTTPClient()
	resp, err := client.Do(req)
	if err != nil {
		return response, code, err
	}

	defer resp.Body.Close()

	// read response code
	code = resp.StatusCode

	// read response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, code, err
	}
	// filter response output
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return response, code, err
	}

	return response, code, nil
}

// UpdateLabs is a function to update labs
func (server CustomServer) UpdateLabs(labId int, labName string) (response map[string]string, code int, err error) {
	// create json data
	jsonData, err := json.Marshal(map[string]interface{}{"name": labName})
	if err != nil {
		return response, code, err
	}

	// create request
	req, err := http.NewRequest("PUT", server.Host+"/v1/admin/labs", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+server.Token)

	// convert int to string
	labIdStr := strconv.Itoa(labId)

	// add query
	q := req.URL.Query()  // Get a copy of the query values.
	q.Add("id", labIdStr) // Add a new value to the set.
	req.URL.RawQuery = q.Encode()

	// send request
	client := getHTTPClient()
	resp, err := client.Do(req)
	if err != nil {
		return response, code, err
	}

	defer resp.Body.Close()

	// read response code
	code = resp.StatusCode

	// read response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, code, err
	}
	// filter response output
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return response, code, err
	}

	return response, code, nil
}

// DeleteLabs is a function to delete labs
func (server CustomServer) DeleteLabs(labId int) (response map[string]string, code int, err error) {
	// create request
	req, err := http.NewRequest("DELETE", server.Host+"/v1/admin/labs", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+server.Token)

	// convert int to string
	labIdStr := strconv.Itoa(labId)

	// add query
	q := req.URL.Query()  // Get a copy of the query values.
	q.Add("id", labIdStr) // Add a new value to the set.
	req.URL.RawQuery = q.Encode()

	// send request
	client := getHTTPClient()
	resp, err := client.Do(req)
	if err != nil {
		return response, code, err
	}

	defer resp.Body.Close()

	// read response code
	code = resp.StatusCode

	// read response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, code, err
	}
	// filter response output
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return response, code, err
	}

	return response, code, nil
}

// CreateCourse is a function to create new course
func (server CustomServer) CreateCourse(courseName string) (response CreateCourseResponse, code int, err error) {
	// create json data
	jsonData, err := json.Marshal(map[string]string{"name": courseName})
	if err != nil {
		return response, code, err
	}

	// create request
	req, err := http.NewRequest("POST", server.Host+"/v1/admin/course", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+server.Token)

	// send request
	client := getHTTPClient()
	resp, err := client.Do(req)
	if err != nil {
		return response, code, err
	}

	defer resp.Body.Close()

	// read response code
	code = resp.StatusCode

	// read response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, code, err
	}
	// filter response output
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return response, code, err
	}

	return response, code, nil
}

// UpdateCourse is a function to update course
func (server CustomServer) UpdateCourse(courseId int, courseName string) (response map[string]string, code int, err error) {
	// create json data
	jsonData, err := json.Marshal(map[string]string{"name": courseName})
	if err != nil {
		return response, code, err
	}

	// create request
	req, err := http.NewRequest("PUT", server.Host+"/v1/admin/course", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+server.Token)

	// convert int to string
	courseIdStr := strconv.Itoa(courseId)

	// add query
	q := req.URL.Query()     // Get a copy of the query values.
	q.Add("id", courseIdStr) // Add a new value to the set.
	req.URL.RawQuery = q.Encode()

	// send request
	client := getHTTPClient()
	resp, err := client.Do(req)
	if err != nil {
		return response, code, err
	}

	defer resp.Body.Close()

	// read response code
	code = resp.StatusCode

	// read response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, code, err
	}
	// filter response output
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return response, code, err
	}

	return response, code, nil
}

// DeleteCourse is a function to delete course
func (server CustomServer) DeleteCourse(courseId int) (response map[string]string, code int, err error) {
	// create request
	req, err := http.NewRequest("DELETE", server.Host+"/v1/admin/course", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+server.Token)

	// convert int to string
	courseIdStr := strconv.Itoa(courseId)

	// add query
	q := req.URL.Query()     // Get a copy of the query values.
	q.Add("id", courseIdStr) // Add a new value to the set.
	req.URL.RawQuery = q.Encode()

	// send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return response, code, err
	}

	defer resp.Body.Close()

	// read response code
	code = resp.StatusCode

	// read response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, code, err
	}
	// filter response output
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return response, code, err
	}

	return response, code, nil
}

type CreateUserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Message  string `json:"message"`
}

type User struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
	ClassID  int    `json:"class_id"`
}

// CreateUser is a function to create new user
func (server CustomServer) CreateUser(user User) (response CreateUserResponse, code int, err error) {
	// create json data
	jsonData, err := json.Marshal(user)
	if err != nil {
		return response, code, err
	}

	// create request
	req, err := http.NewRequest("POST", server.Host+"/v1/admin/user", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+server.Token)

	// send request
	client := getHTTPClient()
	resp, err := client.Do(req)
	if err != nil {
		return response, code, err
	}

	defer resp.Body.Close()

	// read response code
	code = resp.StatusCode

	// read response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, code, err
	}
	// filter response output
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return response, code, err
	}

	return response, code, nil
}

// DeleteUser is a function to delete user
func (server CustomServer) DeleteUser(userId int) (response map[string]string, code int, err error) {
	// create request
	req, err := http.NewRequest("DELETE", server.Host+"/v1/admin/user", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+server.Token)

	// convert int to string
	userIdStr := strconv.Itoa(userId)

	// add query
	q := req.URL.Query()   // Get a copy of the query values.
	q.Add("id", userIdStr) // Add a new value to the set.
	req.URL.RawQuery = q.Encode()

	// send request
	client := getHTTPClient()
	resp, err := client.Do(req)
	if err != nil {
		return response, code, err
	}

	defer resp.Body.Close()

	// read response code
	code = resp.StatusCode

	// read response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, code, err
	}
	// filter response output
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return response, code, err
	}

	return response, code, nil
}

type ExportResponse struct {
	Class   string    `json:"class"`
	Course  string    `json:"course"`
	Date    string    `json:"date"`
	Reports []Reports `json:"reports"`
	Time    string    `json:"time"`
}
type Reports struct {
	Name     string      `json:"name"`
	Username string      `json:"username"`
	Scores   []ScoreLabs `json:"scores"`
	Average  float64     `json:"average"`
	Total    int         `json:"total"`
}

type ScoreLabs struct {
	LabName string `json:"lab_name"`
	Score   int    `json:"score"`
	ID      int    `json:"id"`
}

// ExportScores create api client to handle endpoint /export
func (server CustomServer) ExportScores(courseId int, classId int) (reports ExportResponse, code int, err error) {
	// create request
	req, err := http.NewRequest("GET", server.Host+"/v1/admin/export", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+server.Token)

	// convert int to string
	courseIdStr := strconv.Itoa(courseId)
	classIdStr := strconv.Itoa(classId)

	// add query
	q := req.URL.Query()            // Get a copy of the query values.
	q.Add("course_id", courseIdStr) // Add a new value to the set.
	q.Add("class_id", classIdStr)   // Add a new value to the set.
	req.URL.RawQuery = q.Encode()

	// send request
	client := getHTTPClient()
	resp, err := client.Do(req)
	if err != nil {
		return reports, code, err
	}

	defer resp.Body.Close()

	// read body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return reports, code, err
	}
	// filter response output
	err = json.Unmarshal(bodyBytes, &reports)
	if err != nil {
		return reports, code, err
	}

	code = resp.StatusCode

	return reports, code, nil
}
