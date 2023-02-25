//go:build admin

package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
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
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return response, -1, err
	}

	defer resp.Body.Close()

	// read response code
	code = resp.StatusCode

	// read response body
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
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
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return response, -1, err
	}

	defer resp.Body.Close()

	// read response code
	code = resp.StatusCode

	// read response body
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	// filter response output
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return response, code, err
	}

	return response, code, nil
}

// CreateClass is a function to create new class
func (server CustomServer) CreateClass(className string) (response map[string]string, code int, err error) {
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
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return response, code, err
	}

	defer resp.Body.Close()

	// read response code
	code = resp.StatusCode

	// read response body
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	// filter response output
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return response, code, err
	}

	return response, code, nil
}

// CreateLabs is a function to create new lab
func (server CustomServer) CreateLabs(labName string, courseId int) (response map[string]string, code int, err error) {
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
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return response, code, err
	}

	defer resp.Body.Close()

	// read response code
	code = resp.StatusCode

	// read response body
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	// filter response output
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return response, code, err
	}

	return response, code, nil
}

// CreateCourse is a function to create new course
func (server CustomServer) CreateCourse(courseName string) (response map[string]string, code int, err error) {
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
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return response, code, err
	}

	defer resp.Body.Close()

	// read response code
	code = resp.StatusCode

	// read response body
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	// filter response output
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return response, code, err
	}

	return response, code, nil
}
