//go:build admin

package api

import (
	"testing"
	"time"
)

// create server
var server = CustomServer{
	Host:  "http://localhost:9898",
	Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Nzc2MTgwMzAsInN1YiI6IjMifQ.Wfig6ZiM3weN3Rn95r95Q5vhuv1n1X8gcnWojIsq9qo",
}

var (
	classId  int
	labsId   int
	courseId int
	userId   int
)

// create unit test for cuntion GetUsers
func TestGetUsers(t *testing.T) {
	t.Log("using server", server.Host)
	t.Log("using token", server.Token)
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
	if len(response.Classes) == 0 {
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
	if response.Name != "SIJA 11" {
		t.Error("Class name is not correct")
	}

	// save class id
	classId = response.ID
}

// create unit test for function CreateCourse
func TestCreateCourse(t *testing.T) {
	// call function
	response, code, err := server.CreateCourse("Course Test")
	if err != nil {
		t.Error(err, code)
	}

	// check response
	if response.Name != "Course Test" {
		t.Error("Course name is not correct")
	}

	// check response code
	if code != 201 {
		t.Error("Response code is not correct", code)
	}
	t.Log(response)

	// save course id
	courseId = response.ID
}

// create unit test for function CreateLabs
func TestCreateLabs(t *testing.T) {
	// call function
	response, code, err := server.CreateLabs(2, "linuxfund-1-test2")
	if err != nil {
		t.Error(err, code)
	}

	// check response
	if response.Name != "linuxfund-1-test2" {
		t.Error("CaseLabs name is not correct")
	}

	// check response code
	if code != 201 {
		t.Error("Response code is not correct", code)
	}

	// save labs id
	labsId = response.ID
	t.Log("labs id", labsId)
}

// create unit test for function CreateUser
func TestCreateUser(t *testing.T) {
	wrongUserData := User{
		Name:     "User Test",
		Username: "user-test",
		Password: "user-test",
		ClassID:  classId,
	}

	validUserData := User{
		Name:     "User Test",
		Username: "userTest902",
		Password: "user-test",
		ClassID:  classId,
	}

	// call function
	response, code, err := server.CreateUser(wrongUserData)
	if err != nil {
		t.Log(err, code)
	}

	// check response code
	if code != 400 {
		t.Log("Response code is not correct", code)
	}
	t.Log(response)

	// call function
	response, code, err = server.CreateUser(validUserData)
	if err != nil {
		t.Error(err, code)
	}

	// check response
	if response.Name != "User Test" || response.Username != "userTest902" {
		t.Error("User name is not correct")
	}

	// check response code
	if code != 201 {
		t.Error("Response code is not correct", code)
	}
	t.Log(response)

	// save user id
	userId = response.ID
}

// create unit test for function UpdateLabs
func TestUpdateLabs(t *testing.T) {
	// call function
	response, code, err := server.UpdateLabs(labsId, "linuxfund-1-test3")
	if err != nil {
		t.Error(err, code)
	}

	// check response
	if response["lab"] != "linuxfund-1-test3" {
		t.Error("CaseLabs name is not correct")
	}

	// check response code
	if code != 200 {
		t.Error("Response code is not correct", code)
	}
}

// create unit test for function UpdateClass
func TestUpdateClass(t *testing.T) {
	// call function
	response, code, err := server.UpdateClass(classId, "SIJA 100")
	if err != nil {
		t.Error(err, code)
	}

	// check response
	if response["class"] != "SIJA 100" {
		t.Error("Class name is not correct")
	}

	// check response code
	if code != 200 {
		t.Error("Response code is not correct", code)
	}

	t.Log(response)
}

// create unit test for function UpdateCourse
func TestUpdateCourse(t *testing.T) {
	// call function
	response, code, err := server.UpdateCourse(courseId, "Linux Fundamental 2")
	if err != nil {
		t.Error(err, code)
	}

	// check response
	if response["course"] != "Linux Fundamental 2" {
		t.Error("Course name is not correct")
	}

	// check response code
	if code != 200 {
		t.Error("Response code is not correct", code)
	}
}

// create unit test for function DeleteLabs
func TestDeleteLabs(t *testing.T) {
	// call function
	response, code, err := server.DeleteLabs(labsId)
	if err != nil {
		t.Error(err, code)
	}

	if response["message"] == "" {
		t.Error("Message is empty")
	}
	t.Log(response)

	// check response code
	if code != 200 {
		t.Error("Response code is not correct", code)
	}
}

// create unit test for function DeleteCourse
func TestDeleteCourse(t *testing.T) {
	// call function
	response, code, err := server.DeleteCourse(courseId)
	if err != nil {
		t.Error(err, code)
	}

	if response["message"] == "" {
		t.Error("Message is empty")
	}
	t.Log(response)

	// check response code
	if code != 200 {
		t.Error("Response code is not correct", code)
	}
}

// create unit test for function DeleteUser
func TestDeleteUser(t *testing.T) {
	// call function
	response, code, err := server.DeleteUser(userId)
	if err != nil {
		t.Error(err, code)
	}

	if response["message"] == "" {
		t.Error("Message is empty")
	}
	t.Log(response)

	// check response code
	if code != 200 {
		t.Error("Response code is not correct", code)
	}
}

// create unit test for function DeleteClass
func TestDeleteClass(t *testing.T) {
	// call function
	response, code, err := server.DeleteClass(classId)
	if err != nil {
		t.Error(err, code)
	}

	if response["message"] == "" {
		t.Error("Message is empty")
	}
	t.Log(response)

	// check response code
	if code != 200 {
		t.Error("Response code is not correct", code)
	}
}

// create unit test for function ExportCourse
func TestExportScores(t *testing.T) {
	// call function
	response, code, err := server.ExportScores(2, 1)

	if err != nil {
		t.Error(err, code)
	}

	// check response code
	if code != 200 {
		t.Error("Response code is not correct", code)
	}

	// check response
	if response.Course != "linuxfund" || response.Class != "SIJA 10" {
		t.Error("Response class or Course is not correct")
	}

	// check response
	if response.Date != time.Now().Format("2006-01-02") || response.Time != time.Now().Format("15:04:05") {
		t.Error("Response date is not correct")
	}

	// call function
	users, _, err := server.GetUsers("1")
	if err != nil {
		t.Error(err, code)
	}

	// check response
	if len(response.Reports) != len(users.Students) {
		t.Error("Response length is not correct", len(response.Reports), len(users.Students))
	}

}
