//go:build admin

package api

import (
	"os"
	"testing"
)

// create server
var server = CustomServer{
	Host:  "http://localhost:9898",
	Token: os.Getenv("MOCK_ADMIN_TOKEN"),
	//Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Nzc0MDMzMDUsInN1YiI6IjIifQ.ecywt4lxCZ_3P8QryPybhLj5jEeaXQwtMsUXPJqkxkY",
}

// create unit test for cuntion GetUsers
func TestGetUsers(t *testing.T) {
	// call function
	response, code, err := server.GetUsers("1")
	if err != nil {
		t.Error(err, code)
	}

	// check response
	if response.Class != "SIJA 10" {
		t.Error("Class name is not correct")
	}

	// check total student
	if len(response.Students) != 3 {
		t.Error("Total student is not correct")
	}
}

// create unit test for cuntion GetClasses
func TestGetClasses(t *testing.T) {
	// call function
	response, code, err := server.GetClasses()
	if err != nil {
		t.Error(err, code)
	}

	// check total class
	if len(response.Classes) != 3 {
		t.Error("Total class is not correct", len(response.Classes))
	}
}

// create unit test for function CreateClass
func TestCreateClass(t *testing.T) {
	// call function
	response, code, err := server.CreateClass("SIJA 11")
	if err != nil {
		t.Error(err, code)
	}

	// check response
	if response["class"] != "SIJA 11" {
		t.Error("Class name is not correct")
	}

	// check response code
	if code != 201 {
		t.Error("Response code is not correct", code)
	}

}

// create unit test for function CreateLabs
func TestCreateLabs(t *testing.T) {
	// call function
	response, code, err := server.CreateLabs("linuxfund-1-test2", 2)
	if err != nil {
		t.Error(err, code)
	}

	// check response
	if response["lab"] != "linuxfund-1-test2" {
		t.Error("Labs name is not correct")
	}

	// check response code
	if code != 201 {
		t.Error("Response code is not correct", code)
	}
}

// create unit test for function CreateCourse
func TestCreateCourse(t *testing.T) {
	// call function
	response, code, err := server.CreateCourse("Linux Fundamental")
	if err != nil {
		t.Error(err, code)
	}

	// check response
	if response["course"] != "Linux Fundamental" {
		t.Error("Course name is not correct")
	}

	// check response code
	if code != 201 {
		t.Error("Response code is not correct", code)
	}
}
