//go:build admin

package cmd

import (
	"bufio"
	"encoding/csv"
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

func (r Runner) export(courseId int, classId int) {
	logger.Console("exporting to csv").Start()

	// get data export
	reports, code, err := serverAPI.ExportScores(courseId, classId)

	if code != 200 {
		logger.Console("something went wrong!").Error()
	}

	// Create the CSV file
	file, err := os.Create(fmt.Sprintf("%s-%s_%s.csv", reports.Class, reports.Course, reports.Date))
	if err != nil {
		logger.Console("cant create csv file").Error()
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// getting list of labs name
	labs, code, err := serverAPI.GetLabs(courseId)
	if err != nil || code != 200 {
		logger.Console("cant get list of name").Error()
		fmt.Println(err, code)
		os.Exit(1)
	}

	// Write the header row
	header := []string{"Name", "Username"}
	for _, lab := range labs.Labs {
		header = append(header, lab.Lab)
	}
	header = append(header, "Total", "Average")
	err = writer.Write(header)
	if err != nil {
		logger.Console("cant write to csv").Error()
	}

	// Write the data rows
	for _, student := range reports.Reports {
		data := []string{student.Name, student.Username}
		for _, score := range student.Scores {
			scores := strconv.Itoa(score.Score)
			data = append(data, scores)
		}
		data = append(data, strconv.Itoa(student.Total), strconv.FormatFloat(student.Average, 'f', 2, 64))
		err = writer.Write(data)
		if err != nil {
			logger.Console("cant write to csv").Error()
		}
	}
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
		Short: "admin is manage shopiea resources",
		Long:  `this required admin role, admin is manage shopiea resources such as class, course, student, etc`,
	}

	getCmd := &cobra.Command{
		Use:   "get",
		Short: "print available resources",
		Long:  `print available resources such as class, course, student, etc`,
	}

	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "update available resources",
		Long:  "update resources such as class, course, student, etc",
	}

	removeCmd := &cobra.Command{
		Use:   "remove",
		Short: "remove resources",
		Long:  "remove resources such as class, course, student, etc",
	}

	createCmd := &cobra.Command{
		Use:   "create",
		Short: "create resources",
		Long:  "create resources such as class, course, student, etc",
	}

	getClassCmd := &cobra.Command{
		Use:   "class",
		Short: "get class",
		Long: `print available class
format [id] [name]`,
		Run: func(cmd *cobra.Command, args []string) {
			runner.getClass()
		},
	}

	getCoursesCmd := &cobra.Command{
		Use:   "courses",
		Short: "get courses",
		Long: `print available courses
format [id] [name]`,
		Run: func(cmd *cobra.Command, args []string) {
			runner.getCourses()
		},
	}

	getLabsCmd := &cobra.Command{
		Use:   "labs [COURSE ID]",
		Short: "get labs",
		Long:  "get labs",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			courseId, err := strconv.Atoi(args[0])
			if err != nil {
				logger.Console("Error: " + err.Error()).Error()
			}
			runner.getLabs(courseId)
		},
	}

	getStudentsCmd := &cobra.Command{
		Use:   "students [CLASS ID]",
		Short: "get students",
		Long:  "get students",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runner.getStudents(args[0])
		},
	}

	updateClassCmd := &cobra.Command{
		Use:   "class [CLASS ID] [CLASS NAME]",
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
		Use:   "course [COURSE ID] [COURSE NAME]",
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
		Use:   "lab [LAB ID] [LAB NAME]",
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
		Use:   "class [CLASS NAME]",
		Short: "create class",
		Long:  "create class",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runner.createClass(args[0])
		},
	}

	createCourseCmd := &cobra.Command{
		Use:   "course [COURSE NAME]",
		Short: "create course",
		Long:  "create course",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runner.createCourse(args[0])
		},
	}

	createLabCmd := &cobra.Command{
		Use:   "lab [COURSE ID] [LAB NAME]",
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
		Use:   "student [USERNAME]",
		Short: "create user",
		Long:  "create user",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runner.createStudent(args[0])
		},
	}

	removeStudentCmd := &cobra.Command{
		Use:   "student [USER ID]",
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
		Use:   "class [CLASS ID]",
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
		Use:   "course [COURSE ID]",
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
		Use:   "lab [LAB ID]",
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

	exportCmd := &cobra.Command{
		Use:   "export [COURSE ID] [CLASS ID]",
		Short: "export data",
		Long:  "export data",
		Run: func(cmd *cobra.Command, args []string) {
			// convert string to int
			courseId, err := strconv.Atoi(args[0])
			if err != nil {
				logger.Console("Error: " + err.Error()).Error()
			}

			classId, err := strconv.Atoi(args[1])
			if err != nil {
				logger.Console("Error: " + err.Error()).Error()
			}
			runner.export(courseId, classId)
		},
	}

	getCmd.AddCommand(getClassCmd, getCoursesCmd, getLabsCmd, getStudentsCmd)
	updateCmd.AddCommand(updateClassCmd, updateCourseCmd, updateLabCmd)
	createCmd.AddCommand(createClassCmd, createCourseCmd, createLabCmd, createStudentCmd)
	removeCmd.AddCommand(removeClassCmd, removeCourseCmd, removeLabCmd, removeStudentCmd)
	cmd.AddCommand(getCmd, updateCmd, removeCmd, createCmd, exportCmd)
	return
}

func init() {
	serverAPI = api.CustomServer{
		Host:  runner.Config.GetString("auths.host"),
		Token: runner.Config.GetString("auths.token"),
	}
	philoCmd.AddCommand(adminCommand())
}
