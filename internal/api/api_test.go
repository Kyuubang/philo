package api

import (
	"testing"
	"time"
)

// create server
var serverAPI = CustomServer{
	Host: "http://localhost:9898",
}

// unit test for GetConfig
func TestGetConfig(t *testing.T) {
	t.Log("using server", serverAPI.Host)
	t.Log("using token", serverAPI.Token)

	t.Log("Test Get Config")
	// call function
	response, code, err := serverAPI.GetConfig()
	if err != nil {
		t.Error(err, code)
	}

	t.Log("test response date")
	// check response
	if response["date"] != time.Now().Format(time.RFC822) {
		t.Error("Date is not correct")
	}

	defer func() {
		if r := recover(); r != nil {
			t.Error("Panic error")
		}
	}()

	t.Log("convertion to string test")
	// convert to string
	config := response["config"].(map[string]interface{})

	t.Log("check case branch")
	if config["case_branch"] == "" {
		t.Error("Case branch is not correct")
	}

	if config["case_branch"] != "master" {
		t.Error("Case branch is not correct")
	}
}

// unit test for GetCourses
func TestGetCourses(t *testing.T) {

	t.Log("Test Get Courses")
	// call function
	response, code, err := serverAPI.GetCourses()
	if err != nil {
		t.Error(err, code)
	}

	t.Log("check total courses")
	// check total courses
	if len(response.Courses) == 0 {
		t.Error("Total courses is not correct")
	}
}

// unit test for GetLabs
func TestGetLabs(t *testing.T) {

	t.Log("Test Get CaseLabs")
	// call function
	response, code, err := serverAPI.GetLabs(2)
	if err != nil {
		t.Error(err, code)
	}

	t.Log("check total labs")
	// check total labs
	if len(response.Labs) == 0 {
		t.Error("Total labs is not correct")
	}
}
