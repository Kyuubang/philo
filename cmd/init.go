package cmd

import (
	"github.com/Kyuubang/philo/internal/api"
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
func (r Runner) initRun(lab bool, infraRepo string, caseRepo string) {
	// initialize lab environment
	if lab {
		if os.Getuid() == 0 {
			logger.Console("You must run initiliaze without root user").Warn()
			os.Exit(1)
		}
		_, err := os.Stat(path.Join(r.ConfigPath, "vagrant"))
		if os.IsNotExist(err) {
			err = os.Mkdir(path.Join(r.ConfigPath, "vagrant"), os.FileMode(0766))
			if err != nil {
				logger.Console("unable to create vagrant directory").Error()
			}
			os.Exit(1)
		}
		if err = api.DownloadRepo(infraRepo, path.Join(r.ConfigPath, "vagrant")); err != nil {
			r.Log.MainLog.Error().Msgf("unable to download infra repo: %s", err)
			logger.Console("unable to download infra repo").Error()
			os.Exit(1)
		}
		return
	}

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
		Short: "initialize your philo environment",
		Long: `initialize your philo environment, it would be creating default necessary directory, 
download default infra and test case repo, and set default configuration`,
		Run: func(cmd *cobra.Command, args []string) {
			var lab, _ = cmd.Flags().GetBool("lab")
			var infraRepo, _ = cmd.Flags().GetString("infra-repo")
			var testCaseRepo, _ = cmd.Flags().GetString("test-case-repo")
			runner.initRun(lab, infraRepo, testCaseRepo)
		},
	}

	var infraRepo string
	var testCaseRepo string

	// create flag to specify infra repo
	cmd.Flags().StringVarP(&infraRepo, "infra-repo", "i", "", "Default infra repo to use")
	// create flag to specify test case repo
	cmd.Flags().StringVarP(&testCaseRepo, "test-case-repo", "t", "", "Default test case repo to use")
	// create flag to used for lab
	cmd.Flags().BoolP("lab", "l", false, "Initialize for lab environment")

	return cmd
}
