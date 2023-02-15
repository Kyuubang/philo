package cmd

import (
	"fmt"
	"github.com/Kyuubang/philo/internal/api"
	"github.com/Kyuubang/philo/internal/utils/remote"
	"github.com/Kyuubang/philo/logger"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
	"os"
)

func score(point int, total int) int {
	if total == 0 {
		return 0
	} else if point > total {
		return 0
	} else if total == point {
		return 100
	} else {
		score := (100 / total) * point
		return score
	}
}

func (r Runner) scoreCheck(labName string) {

	logger.Console("Checking score").Start()

	result, code := api.GetCase(r.Config.GetString("repo"), "master", labName)

	switch code {
	case 403:
		logger.Console("Forbidden!").Error()
		os.Exit(1)
	case 404:
		logger.Console("Lab not found!").Error()
		os.Exit(1)
	case 200:
		logger.Console("Lab found!").Success()
	default:
		logger.Console("Unknown error").Error()
		os.Exit(1)
	}
	// setup remote command setup
	VMSSHConfig, err := getVMSSHConfig(labName)
	if err != nil {
		logger.Console("Error getting VM SSH config").Error()
		os.Exit(1)
	}

	var checkPoint int
	var color = logger.Color{}

	for _, grade := range result.Grade {
		logger.Console("Checking " + grade.Name).Test()

		// setup remote command setup
		sshConfig := &ssh.ClientConfig{
			User: VMSSHConfig[grade.On].User,
			Auth: []ssh.AuthMethod{
				remote.PublicKeyFile(VMSSHConfig[grade.On].IdentityFile),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}

		client := remote.SSHClient{
			Config: sshConfig,
			Host:   VMSSHConfig[grade.On].HostName,
			Port:   VMSSHConfig[grade.On].Port,
		}

		out, err := client.RunRemoteCommand(grade.Script)
		if err != nil {
			logger.Console("Error running remote command").Error()

		}

		if grade.Expect == out.StdOut {
			checkPoint++
			fmt.Println(fmt.Sprintf(" [%s]", color.Green("✓")))
		} else {
			r.Log.MainLog.Error().Str("lab", labName).Str("case", grade.Name)
			fmt.Println(fmt.Sprintf(" [%s]", color.Red("x")))
		}
	}

	logger.Console(fmt.Sprintf("Score: %d/%d", score(checkPoint, len(result.Grade)), 100)).Info()

	logger.Console("Score check success").Success()
}

func (r Runner) scoreView(labName string) {
	logger.Console("Viewing score").Start()

	logger.Console("Score view success").Success()
}

func scoreCommand() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "score",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	}

	cmd.AddCommand(
		&cobra.Command{
			Use:   "check",
			Short: "A brief description of your command",
			Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
			Args: cobra.MinimumNArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				runner.scoreCheck(args[0])
			},
		},
		&cobra.Command{
			Use:   "view",
			Short: "A brief description of your command",
			Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
			Args: cobra.MinimumNArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				runner.scoreView(args[0])
			},
		})

	return cmd
}
