package cmd

import (
	"fmt"
	"github.com/Kyuubang/philo/internal/api"
	"github.com/Kyuubang/philo/logger"
	"github.com/spf13/cobra"
)

func (r Runner) courseList(registered bool) {
	logger.Console("courses list").Start()

	if registered {
		// list all course with api.GetCourse()
		cServer := api.CustomServer{
			Host:  r.Config.GetString("auths.host"),
			Token: r.Config.GetString("auths.token"),
		}

		courses, _, err := cServer.GetCourses()
		if err != nil {
			logger.Console("Error: " + err.Error()).Error()
		}

		for _, course := range courses.Courses {
			logger.Console(fmt.Sprintf("%d -> %s", course.ID, course.Name)).Info()
		}
	}

	logger.Console("courses list").Success()
}

func courseCommand() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "course",
		Short: "course is a command to manage course",
	}

	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "list is a command to list available course",
		Run: func(cmd *cobra.Command, args []string) {
			registered, err := cmd.Flags().GetBool("registered")
			if err != nil {
				logger.Console("Error: " + err.Error()).Error()
			}
			runner.courseList(registered)
		},
	}

	listCmd.Flags().BoolP("registered", "r", false, "list registered courses")

	cmd.AddCommand(listCmd)

	return
}
