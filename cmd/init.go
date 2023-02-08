package cmd

import (
	"github.com/Kyuubang/philo/logger"
	"github.com/spf13/cobra"
	"os"
	"path"
)

const (
	logPath = "/var/log/philo"
	//fileLogDefault = "philo.log"
)

// Run main actions handler for init command
// it will be creating log path -> /var/log/philo with rwx perm
// and chowner to 1000 uid's user
func (r Runner) initRun() {
	if os.Getuid() != 0 {
		logger.Console("You must run initiliaze with root user").Warn()
		os.Exit(1)
	}

	// create logging path -> /var/log/philo
	//fmt.Println("Run Initilize Configuration ", "info")
	logger.Console("Run Initilize Configuration ").Start()
	_, err := os.Stat(logPath)
	if os.IsNotExist(err) {
		err = os.Mkdir(logPath, os.FileMode(0766))
		if err != nil {
			logger.Console("unable to create logging directory").Error()
		}
	}

	// create logging paths labs -> /var/log/philo/labs
	_, err = os.Stat(path.Join(logPath, "labs"))
	if os.IsNotExist(err) {
		err = os.Mkdir(path.Join(logPath, "labs"), os.FileMode(0766))
		if err != nil {
			logger.Console("unable to create labs logging directory").Error()
		}
	}

	// setting permission and owner
	err = os.Chown(logPath+"/labs", 1000, 0)
	if err != nil {
		logger.Console("unable to set owner labs logging directory").Error()
	}
	err = os.Chown(logPath, 1000, 0)
	if err != nil {
		logger.Console("unable to set owner logging directory ").Error()
	}
	logger.Console("Initilized Success").Success()
}

func initCommand() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "init",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			runner.initRun()
		},
	}

	return cmd
}
