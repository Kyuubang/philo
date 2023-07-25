package cmd

import (
	"fmt"
	"github.com/Kyuubang/philo/internal/api"
	"github.com/Kyuubang/philo/internal/utils/remote"
	"github.com/Kyuubang/philo/logger"
	"github.com/bmatcuk/go-vagrant"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"os"
)

var (
	file     string
	vagrants bool
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
	var VMSSHConfig map[string]vagrant.SSHConfig
	var caseData api.CaseData
	var push bool
	var vmSpec string

	// check if --file used
	if file != "" {
		logger.Console("Checking score from file " + file).Info()

		// convert to file to struct
		// Check if file exists
		if _, err := os.Stat(file); err != nil {
			if os.IsNotExist(err) {
				logger.Console("File not found").Error()
				os.Exit(1)
			} else {
				logger.Console("Error opening file").Error()
				os.Exit(1)
			}
		}

		content, err := ioutil.ReadFile(file)
		if err != nil {
			logger.Console("Error reading file").Error()
			os.Exit(1)
		}

		// Parse file
		parserCase, err := api.FileCaseParser(content)
		if err != nil {
			logger.Console("Error parsing file").Error()
			fmt.Println(err)
			os.Exit(1)
		}

		caseData = parserCase
		push = false

	} else {
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

		caseData = result
		push = true
	}

	if vagrants {
		vmSpec = caseData.Spec

		// setup remote command setup
		SSHConfig, err := getVMSSHConfig(vmSpec)
		if err != nil {
			logger.Console("Error getting VM SSH config").Error()
			os.Exit(1)
		}

		VMSSHConfig = SSHConfig
	} else {
		// check host
		var hosts = r.Config.GetStringMap("hosts")

		// Create an output map
		VMSSHConfig = make(map[string]vagrant.SSHConfig)

		// Iterate over the input map
		for key, value := range hosts {
			// Cast the value to MyStruct
			values := value.(map[string]interface{})
			if values["identity_file"] == nil {
				values["identity_file"] = ""
			}
			if values["password"] == nil {
				values["password"] = ""
			}
			defer func() {
				if r := recover(); r != nil {
					logger.Console("Error parsing host config").Error()
					fmt.Println(r)
					os.Exit(1)
				}
			}()

			host := vagrant.SSHConfig{
				HostName:               values["hostname"].(string),
				User:                   values["user"].(string),
				Port:                   int(values["port"].(float64)),
				IdentityFile:           values["identity_file"].(string),
				PasswordAuthentication: values["password"].(string),
			}

			// Add the MyStruct to the output map
			VMSSHConfig[key] = host
		}
	}

	var checkPoint int
	var color = logger.Color{}

	if labName != caseData.Slug {
		logger.Console("Lab name not match").Error()
		os.Exit(1)
	}

	for _, grade := range caseData.Grade {
		logger.Console("Checking " + grade.Name).Test()

		privKey, err := remote.PublicKeyFile(VMSSHConfig[grade.On].IdentityFile)
		if err != nil {
			r.Log.MainLog.Error().Str("host", grade.On).Msg(err.Error())
		}

		// setup remote command setup
		sshConfig := &ssh.ClientConfig{
			User: VMSSHConfig[grade.On].User,
			Auth: []ssh.AuthMethod{
				ssh.Password(VMSSHConfig[grade.On].PasswordAuthentication),
				privKey,
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

	logger.Console(fmt.Sprintf("Score: %d/%d", score(checkPoint, len(caseData.Grade)), 100)).Info()

	if push {
		apis := api.CustomServer{
			Host:  r.Config.GetString("auths.host"),
			Token: r.Config.GetString("auths.token"),
		}

		// push score to api
		scoreData := api.ScoreData{
			Username: r.Config.GetString("auths.username"),
			Lab:      labName,
			Score:    score(checkPoint, len(caseData.Grade)),
		}

		code, err := apis.PushScore(scoreData)
		if err != nil {
			logger.Console("Error pushing score").Error()
			os.Exit(1)
		}

		switch code {
		case 201:
			logger.Console("Score pushed").Success()
		case 202:
			logger.Console("Dont worry we keep higher score").Success()
		default:
			logger.Console("cant push score").Error()
		}
	}
}

func (r Runner) scoreView(labName string) {
	logger.Console("Viewing score " + labName).Start()

	apis := api.CustomServer{
		Host:  r.Config.GetString("auths.host"),
		Token: r.Config.GetString("auths.token"),
	}

	// TODO: catch error within response code

	score, _, err := apis.GetScore(labName)
	if err != nil {
		logger.Console("Error getting score").Error()
		os.Exit(1)
	}

	logger.Console(fmt.Sprintf("Score: %d/%d", score.Score, 100)).Info()

	logger.Console("Score view success").Success()
}

func scoreCommand() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "score",
		Short: "manage score of lab",
		Long:  `manage score of lab you can check score or view score`,
	}

	checkCmd := &cobra.Command{
		Use:   "check [LAB]",
		Short: "run all case and return score",
		Long:  `run all case and return score, it grab case.yaml from repo and run it on vm`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runner.scoreCheck(args[0])
		},
	}

	// add flag --file to check command
	checkCmd.Flags().StringVarP(&file, "file", "f", "", "file to check")
	checkCmd.Flags().BoolVarP(&vagrants, "vagrant", "s", false, "enable managed vagrant ssh config")

	viewCmd := &cobra.Command{
		Use:   "view [LAB]",
		Short: "view score of lab",
		Long:  `view score of lab if you already run check`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runner.scoreView(args[0])
		},
	}

	cmd.AddCommand(checkCmd, viewCmd)

	return cmd
}
