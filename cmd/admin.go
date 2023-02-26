//go:build admin

package cmd

import (
	"bufio"
	"fmt"
	"github.com/Kyuubang/philo/internal/api"
	"github.com/Kyuubang/philo/logger"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"regexp"
	"strconv"
	"strings"
	"syscall"
)

// TODO: create function to provide admin operations

var (
	serverAPI api.CustomServer
)

func (r Runner) getClass() {
	logger.Console("get class").Start()

	// get class with api
	response, code, err := serverAPI.GetClasses()
	if err != nil {
		logger.Console("Error: " + err.Error()).Error()
	}

	if code != 200 {
		logger.Console(fmt.Sprintf("Error : %d", code)).Error()
	}

	for _, class := range response.Classes {
		logger.Console(fmt.Sprintf("%d -> %s", class.ID, class.Name)).Info()
	}

	logger.Console("get class").Success()
}

func (r Runner) getCourses() {
	logger.Console("get courses").Start()

	// get courses with api
	response, code, err := serverAPI.GetCourses()
	if err != nil {
		logger.Console("Error: " + err.Error()).Error()
	}

	if code != 200 {
		logger.Console(fmt.Sprintf("Error : %d", code)).Error()
	}

	for _, course := range response.Courses {
		logger.Console(fmt.Sprintf("%d -> %s", course.ID, course.Name)).Info()
	}

	logger.Console("get courses").Success()
}

func (r Runner) getLabs(courseId int) {
	logger.Console("get labs").Start()

	// get labs with api
	response, code, err := serverAPI.GetLabs(courseId)
	if err != nil {
		logger.Console("Error: " + err.Error()).Error()
	}

	if code != 200 {
		logger.Console(fmt.Sprintf("Error : %d", code)).Error()
	}

	for _, lab := range response.Labs {
		logger.Console(fmt.Sprintf("%d -> %s", lab.ID, lab.Lab)).Info()
	}

	logger.Console("get labs").Success()
}

func (r Runner) getStudents(classId string) {
	logger.Console("get students").Start()

	// get students with api
	response, code, err := serverAPI.GetUsers(classId)
	if err != nil {
		logger.Console("Error: " + err.Error()).Error()
	}

	if code != 200 {
		logger.Console(fmt.Sprintf("Error : %d", code)).Error()
	}

	for _, student := range response.Students {
		logger.Console(fmt.Sprintf("%d -> %s", student.ID, student.Name)).Info()
	}

	logger.Console("get students").Success()
}

func (r Runner) updateClass(classId int, className string) {
	logger.Console("update class").Start()

	// update class with api
	response, code, err := serverAPI.UpdateClass(classId, className)
	if err != nil {
		logger.Console("Error: " + err.Error()).Error()
	}

	if code != 200 {
		logger.Console(fmt.Sprintf("Error : %d", code)).Error()
	}

	logger.Console(response["message"]).Success()
}

func (r Runner) updateCourse(courseId int, courseName string) {
	logger.Console("update course").Start()

	// update course with api
	response, code, err := serverAPI.UpdateCourse(courseId, courseName)
	if err != nil {
		logger.Console("Error: " + err.Error()).Error()
	}

	if code != 200 {
		logger.Console(fmt.Sprintf("Error : %d", code)).Error()
	}

	logger.Console(response["message"]).Success()
}

func (r Runner) updateLab(labId int, labName string) {
	logger.Console("update lab").Start()

	// update lab with api
	response, code, err := serverAPI.UpdateLabs(labId, labName)
	if err != nil {
		logger.Console("Error: " + err.Error()).Error()
	}

	if code != 200 {
		logger.Console(fmt.Sprintf("Error : %d", code)).Error()
	}

	logger.Console(response["message"]).Success()
}

func (r Runner) createClass(className string) {
	logger.Console("create class").Start()

	// create class with api
	response, code, err := serverAPI.CreateClass(className)
	if err != nil {
		logger.Console("Error: " + err.Error()).Error()
	}

	if code != 201 {
		logger.Console(fmt.Sprintf("Error : %d", code)).Error()
	}

	logger.Console(response.Message).Success()
}

func (r Runner) createCourse(courseName string) {
	logger.Console("create course").Start()

	// create course with api
	response, code, err := serverAPI.CreateCourse(courseName)
	if err != nil {
		logger.Console("Error: " + err.Error()).Error()
	}

	if code != 200 {
		logger.Console(fmt.Sprintf("Error : %d", code)).Error()
	}

	logger.Console(response.Message).Success()
}

func (r Runner) createLab(courseId int, labName string) {
	logger.Console("create lab").Start()

	// create lab with api
	response, code, err := serverAPI.CreateLabs(courseId, labName)
	if err != nil {
		logger.Console("Error: " + err.Error()).Error()
	}

	if code != 200 {
		logger.Console(fmt.Sprintf("Error : %d", code)).Error()
	}

	logger.Console(response.Message).Success()
}

func (r Runner) removeClass(classId int) {
	logger.Console("remove class").Start()

	// remove class with api
	response, code, err := serverAPI.DeleteClass(classId)
	if err != nil {
		logger.Console("Error: " + err.Error()).Error()
	}

	if code != 200 {
		logger.Console(fmt.Sprintf("Error : %d", code)).Error()
	}

	logger.Console(response["message"]).Success()
}

func (r Runner) removeCourse(courseId int) {
	logger.Console("remove course").Start()

	// remove course with api
	response, code, err := serverAPI.DeleteCourse(courseId)
	if err != nil {
		logger.Console("Error: " + err.Error()).Error()
	}

	if code != 200 {
		logger.Console(fmt.Sprintf("Error : %d", code)).Error()
	}

	logger.Console(response["message"]).Success()
}

func (r Runner) removeLab(labId int) {
	logger.Console("remove lab").Start()

	// remove lab with api
	response, code, err := serverAPI.DeleteLabs(labId)
	if err != nil {
		logger.Console("Error: " + err.Error()).Error()
	}

	if code != 200 {
		logger.Console(fmt.Sprintf("Error : %d", code)).Error()
	}

	logger.Console(response["message"]).Success()
}

func isValidUsername(s string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9]{1,16}$`)
	return re.MatchString(s)
}

// function to create student
func (r Runner) createStudent(userName string) {
	// check username is valid
	if !isValidUsername(userName) {
		logger.Console("Error: username is must alphanumeric and no longer than 16 char").Error()
		os.Exit(1)
	}

	logger.Console("create student " + userName).Start()

	// create stdin reader for Name input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Name: ")
	studentName, _ := reader.ReadString('\n')
	// remove new line character
	studentName = strings.TrimSuffix(studentName, "\n")

	// create stdin reader for Class input
	reader = bufio.NewReader(os.Stdin)
	fmt.Print("Enter Class: ")
	classIdString, _ := reader.ReadString('\n')
	// remove new line character
	classIdString = strings.TrimSuffix(classIdString, "\n")

	classId, err := strconv.Atoi(classIdString)
	if err != nil {
		logger.Console("Error: " + err.Error()).Error()
	}

	// input password for student with validation
	var studentPassword string
	for {
		fmt.Print("Enter Password: ")
		bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
		studentPassword = string(bytePassword)
		fmt.Println()
		fmt.Print("Confirm Password: ")
		bytePassword, _ = terminal.ReadPassword(int(syscall.Stdin))
		confirmPassword := string(bytePassword)
		fmt.Println()
		if studentPassword == confirmPassword {
			break
		}
		logger.Console("Password does not match").Error()
	}

	// create student struct with User struct
	student := api.User{
		Username: userName,
		Name:     studentName,
		Password: studentPassword,
		ClassID:  classId,
	}

	// create student with api
	response, code, err := serverAPI.CreateUser(student)
	if err != nil {
		logger.Console("Error: " + err.Error()).Error()
	}

	if code != 201 {
		logger.Console(fmt.Sprintf("Error : %d", code)).Error()
	}

	logger.Console(response.Message).Success()
}

func (r Runner) removeStudent(studentId int) {
	logger.Console("remove student").Start()

	// remove student with api
	response, code, err := serverAPI.DeleteUser(studentId)
	if err != nil {
		logger.Console("Error: " + err.Error()).Error()
	}

	if code != 200 {
		logger.Console(fmt.Sprintf("Error : %d", code)).Error()
	}

	logger.Console(response["message"]).Success()
}

func adminCommand() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "admin",
		Short: "config is a command to manage config file",
		Long:  `config is a command to manage config file`,
	}

	getCmd := &cobra.Command{
		Use:   "get",
		Short: "print available resources",
		Long:  `print available resources`,
	}

	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "update available resources",
		Long:  "update available resources",
	}

	removeCmd := &cobra.Command{
		Use:   "remove",
		Short: "remove available resources",
		Long:  "remove available resources",
	}

	createCmd := &cobra.Command{
		Use:   "create",
		Short: "create available resources",
		Long:  "create available resources",
	}

	getClassCmd := &cobra.Command{
		Use:   "class",
		Short: "get class",
		Long:  "get class",
		Run: func(cmd *cobra.Command, args []string) {
			runner.getClass()
		},
	}

	getCoursesCmd := &cobra.Command{
		Use:   "courses",
		Short: "get courses",
		Long:  "get courses",
		Run: func(cmd *cobra.Command, args []string) {
			runner.getCourses()
		},
	}

	getLabsCmd := &cobra.Command{
		Use:   "labs",
		Short: "get labs",
		Long:  "get labs",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			courseId, err := strconv.Atoi(args[0])
			if err != nil {
				logger.Console("Error: " + err.Error()).Error()
			}
			runner.getLabs(courseId)
		},
	}

	getStudentsCmd := &cobra.Command{
		Use:   "students",
		Short: "get students",
		Long:  "get students",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runner.getStudents(args[0])
		},
	}

	updateClassCmd := &cobra.Command{
		Use:   "class",
		Short: "update class",
		Long:  "update class",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			classId, err := strconv.Atoi(args[0])
			if err != nil {
				logger.Console("Error: " + err.Error()).Error()
			}
			runner.updateClass(classId, args[1])
		},
	}

	updateCourseCmd := &cobra.Command{
		Use:   "course",
		Short: "update course",
		Long:  "update course",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			courseId, err := strconv.Atoi(args[0])
			if err != nil {
				logger.Console("Error: " + err.Error()).Error()
			}
			runner.updateCourse(courseId, args[1])
		},
	}

	updateLabCmd := &cobra.Command{
		Use:   "lab",
		Short: "update lab",
		Long:  "update lab",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			labId, err := strconv.Atoi(args[0])
			if err != nil {
				logger.Console("Error: " + err.Error()).Error()
			}
			runner.updateLab(labId, args[1])
		},
	}

	createClassCmd := &cobra.Command{
		Use:   "class",
		Short: "create class",
		Long:  "create class",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runner.createClass(args[0])
		},
	}

	createCourseCmd := &cobra.Command{
		Use:   "course",
		Short: "create course",
		Long:  "create course",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runner.createCourse(args[0])
		},
	}

	createLabCmd := &cobra.Command{
		Use:   "lab",
		Short: "create lab",
		Long:  "create lab",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			courseId, err := strconv.Atoi(args[0])
			if err != nil {
				logger.Console("Error: " + err.Error()).Error()
			}
			runner.createLab(courseId, args[1])
		},
	}

	createStudentCmd := &cobra.Command{
		Use:   "student [username]",
		Short: "create user",
		Long:  "create user",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runner.createStudent(args[0])
		},
	}

	removeStudentCmd := &cobra.Command{
		Use:   "student [userId]",
		Short: "remove user",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			userId, err := strconv.Atoi(args[0])
			if err != nil {
				logger.Console("Error: " + err.Error()).Error()
			}
			runner.removeStudent(userId)
		},
	}

	removeClassCmd := &cobra.Command{
		Use:   "class [classId]",
		Short: "remove class",
		Long:  "remove class",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			classId, err := strconv.Atoi(args[0])
			if err != nil {
				logger.Console("Error: " + err.Error()).Error()
			}
			runner.removeClass(classId)
		},
	}

	removeCourseCmd := &cobra.Command{
		Use:   "course [courseId]",
		Short: "remove course",
		Long:  "remove course",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			courseId, err := strconv.Atoi(args[0])
			if err != nil {
				logger.Console("Error: " + err.Error()).Error()
			}
			runner.removeCourse(courseId)
		},
	}

	removeLabCmd := &cobra.Command{
		Use:   "lab [labId]",
		Short: "remove lab",
		Long:  "remove lab",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			labId, err := strconv.Atoi(args[0])
			if err != nil {
				logger.Console("Error: " + err.Error()).Error()
			}
			runner.removeLab(labId)
		},
	}

	getCmd.AddCommand(getClassCmd, getCoursesCmd, getLabsCmd, getStudentsCmd)
	updateCmd.AddCommand(updateClassCmd, updateCourseCmd, updateLabCmd)
	createCmd.AddCommand(createClassCmd, createCourseCmd, createLabCmd, createStudentCmd)
	removeCmd.AddCommand(removeClassCmd, removeCourseCmd, removeLabCmd, removeStudentCmd)
	cmd.AddCommand(getCmd, updateCmd, removeCmd, createCmd)

	return
}

func init() {
	serverAPI = api.CustomServer{
		Host:  runner.Config.GetString("auths.host"),
		Token: runner.Config.GetString("auths.token"),
	}
	philoCmd.AddCommand(adminCommand())
}
