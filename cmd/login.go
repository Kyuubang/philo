package cmd

import (
	"bufio"
	"fmt"
	"github.com/Kyuubang/philo/internal/api"
	"github.com/Kyuubang/philo/logger"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"os"
	"strings"
	"syscall"
)

func credentials() (string, string, error) {
	reader := bufio.NewReader(os.Stdin)

	// getting input user name
	fmt.Print("Username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		return "", "", err
	}

	// getting password input
	fmt.Print("Password: ")
	bytePassword, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		return "", "", err
	}

	password := string(bytePassword)
	return strings.TrimSpace(username), strings.TrimSpace(password), nil
}

func (r Runner) login(host string) {
	r.Log.MainLog.Info().Msg("Login")
	logger.Console("Login to " + host).Start()

	// get username and password
	user, pass, _ := credentials()
	fmt.Print("\n")

	// login to custom server
	if host != "github.com" {

		access, err := api.Login(host, user, pass)
		if err != nil {
			// check if error is ErrorLoginFailed
			switch err {
			case api.ErrorUserOrPass:
				logger.Console("Username or password is incorrect").Error()
				os.Exit(1)
			case api.ErrorServer:
				logger.Console("Server error").Error()
				os.Exit(1)
			default:
				fmt.Println(err)
				logger.Console("Unknown error").Error()
			}
		}

		// write to config.json
		authsData := map[string]string{
			"host":     host,
			"username": user,
			"token":    access.Token,
		}

		r.Config.Set("auths", authsData)
		err = r.Config.WriteConfig()
		if err != nil {
			logger.Console("Error writing config file").Error()
			os.Exit(1)
		}

	}
}

func loginCommand() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "login [SERVER]",
		Short: "login to server",
		Long:  `login to server, server should be include protocol and port`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				logger.Console("Missing server").Error()
				os.Exit(1)
			}
			runner.login(args[0])
		},
	}

	return
}
