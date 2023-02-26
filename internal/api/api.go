package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CustomServer struct {
	Host  string
	Token string
}

var (
	ErrorLoginFailed = fmt.Errorf("Login failed")
	ErrorUserOrPass  = fmt.Errorf("Username or password is empty")
	ErrorGeneral     = fmt.Errorf("General error")
	ErrorServer      = fmt.Errorf("Server error")
)

// Login create authentication to server
func Login(host string, username string, password string) (response LoginResponse, err error) {
	var data = LoginData{
		Username: username,
		Password: password,
	}

	// create json data
	jsonData, err := json.Marshal(data)
	if err != nil {
		return response, err
	}

	// create request
	req, err := http.Post(host+"/auth/login", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return response, err
	}

	defer req.Body.Close()

	var userr LoginResponse

	// read response body
	bodyBytes, _ := ioutil.ReadAll(req.Body)
	// filter response output
	err = json.Unmarshal(bodyBytes, &userr)
	if err != nil {
		return response, err
	}

	// check http status code
	switch req.StatusCode {
	case 401:
		return response, ErrorUserOrPass
	case 400:
		return response, ErrorGeneral
	case 200:
		return userr, nil
	default:
		return response, ErrorServer
	}
}

func (server CustomServer) GetCourses() (course CoursesResponse, code int, err error) {
	// create request with function getBody
	req, err := http.NewRequest("GET", server.Host+"/v1/course", nil)

	if err != nil {
		return course, code, err
	}

	// set header
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+server.Token)

	// send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return course, code, err
	}

	defer resp.Body.Close()

	// response code
	code = resp.StatusCode

	// read response body
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	// filter response output
	err = json.Unmarshal(bodyBytes, &course)
	if err != nil {
		return course, code, err
	}

	return course, code, nil
}

func (server CustomServer) GetLabs(courseId int) (labs LabsResponse, code int, err error) {
	// create request
	req, err := http.NewRequest("GET", server.Host+"/v1/labs", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+server.Token)

	q := req.URL.Query()                            // Get a copy of the query values.
	q.Add("course_id", fmt.Sprintf("%d", courseId)) // Add a new value to the set.
	req.URL.RawQuery = q.Encode()

	// send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return labs, code, err
	}

	defer resp.Body.Close()

	// response code
	code = resp.StatusCode

	// read response body
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	// filter response output
	err = json.Unmarshal(bodyBytes, &labs)
	if err != nil {
		return labs, code, err
	}

	return labs, code, nil
}

func (server CustomServer) PushScore(data ScoreData) (code int, err error) {
	// create json data
	jsonData, err := json.Marshal(data)
	if err != nil {
		return code, err
	}

	// create request
	req, err := http.NewRequest("POST", server.Host+"/v1/score", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+server.Token)

	// send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return code, err
	}

	defer resp.Body.Close()

	return resp.StatusCode, nil
}

func (server CustomServer) GetScore(labName string) (score GetScoreResponse, code int, err error) {
	// create request
	req, err := http.NewRequest("GET", server.Host+"/v1/score", nil)

	q := req.URL.Query()       // Get a copy of the query values.
	q.Add("lab_name", labName) // Add a new value to the set.
	req.URL.RawQuery = q.Encode()

	if err != nil {
		return score, code, err
	}

	// set header
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+server.Token)

	// send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return score, code, err
	}

	defer resp.Body.Close()

	// read response body
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	// filter response output
	err = json.Unmarshal(bodyBytes, &score)
	if err != nil {
		return score, code, err
	}

	return score, resp.StatusCode, nil
}

// GetConfig to get config from server at /v1/info endpoint
func (server CustomServer) GetConfig() (config map[string]interface{}, code int, err error) {
	// create request
	req, err := http.NewRequest("GET", server.Host+"/v1/info", nil)
	if err != nil {
		return config, -1, err
	}

	// set header
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+server.Token)

	// send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return config, -1, err
	}

	defer resp.Body.Close()

	code = resp.StatusCode

	// read response body
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	// filter response output
	err = json.Unmarshal(bodyBytes, &config)
	if err != nil {
		return config, code, err
	}

	return config, code, nil

}
