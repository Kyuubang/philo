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

type LoginResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}

var (
	ErrorLoginFailed = fmt.Errorf("Login failed")
	ErrorUserOrPass  = fmt.Errorf("Username or password is empty")
	ErrorGeneral     = fmt.Errorf("General error")
	ErrorServer      = fmt.Errorf("Server error")
)

// create function to login to server
func (server CustomServer) Login(username string, password string) (response LoginResponse, err error) {
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
	req, err := http.Post(server.Host+"/auth/login", "application/json", bytes.NewBuffer(jsonData))
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
