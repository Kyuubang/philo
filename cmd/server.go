package cmd

import (
	"bufio"
	"fmt"
	"github.com/Kyuubang/philo/internal/api"
	"github.com/Kyuubang/philo/internal/utils/remote"
	"github.com/Kyuubang/philo/logger"
	"github.com/bmatcuk/go-vagrant"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
	"io"
	"os"
)

// GetSSHConfig returns the SSH config for the given lab
func getVMSSHConfig(vagrantDir string) (sshConfig map[string]vagrant.SSHConfig, err error) {
	client, err := vagrant.NewVagrantClient(vagrantDir)

	if err != nil {
		return nil, err
	}

	vagrantSSHConfig := client.SSHConfig()

	ok := vagrantSSHConfig.Run()
	if ok != nil {
		return nil, ok
	}

	if vagrantSSHConfig.Error != nil {
		return nil, vagrantSSHConfig.ErrorResponse.Error
	}

	response := vagrantSSHConfig.SSHConfigResponse
	sshConfig = response.Configs
	return sshConfig, nil
}

// GetStatus returns the status of the given lab
func getVMStatus(vagrantDir string) (status map[string]string, err error) {
	client, err := vagrant.NewVagrantClient(vagrantDir)
	if err != nil {
		return nil, err
	}

	statusCommand := client.Status()

	ok := statusCommand.Run()
	if ok != nil {
		return nil, ok
	}

	if statusCommand.Error != nil {
		return nil, statusCommand.ErrorResponse.Error
	}

	return statusCommand.StatusResponse.Status, nil
}

func (r Runner) serverShow(labName string) {
	logger.Console("Showing server").Start()

	cases, code := api.GetCase(r.Config.GetString("repo"), "master", labName)
	if code != 200 {
		r.Log.MainLog.Error().Msg("Error getting lab info")
		logger.Console("Error getting lab info").Error()
		os.Exit(1)
	}

	specVM := fmt.Sprintf(r.ConfigPath+"/vagrant/%s", cases.Spec)
	//fmt.Println("Vagrant spec: " + specVM)
	var statusVM, err = getVMStatus(specVM)
	if err != nil {
		r.Log.MainLog.Error().Msg(err.Error())
		os.Exit(1)
	}

	for machine, status := range statusVM {
		fmt.Printf("    %s: %s\n", machine, status)
	}
	logger.Console("Vagrant status success").Success()
}

func (r Runner) serverHalt(labName string) {
	logger.Console("Halting server").Start()

	cases, code := api.GetCase(r.Config.GetString("repo"), "master", labName)
	if code != 200 {
		r.Log.MainLog.Error().Msg("Error getting lab info")
		logger.Console("Error getting lab info").Error()
		os.Exit(1)
	}

	specVM := fmt.Sprintf(r.ConfigPath+"/vagrant/%s", cases.Spec)

	client, err := vagrant.NewVagrantClient(specVM)
	if err != nil {
		r.Log.MainLog.Error().Msg(err.Error())
		os.Exit(1)
	}

	haltCommand := client.Halt()

	ok := haltCommand.Run()
	if ok != nil {
		r.Log.MainLog.Error().Msg(ok.Error())
		logger.Console("Vagrant halt failed").Error()
		os.Exit(1)
	}

	if haltCommand.Error != nil {
		r.Log.MainLog.Error().Msg(haltCommand.ErrorResponse.Error.Error())
		os.Exit(1)
	}

	logger.Console("Vagrant halt success").Success()
}

func (r Runner) serverDestroy(labName string) {
	logger.Console("Destroying server " + labName).Start()

	cases, code := api.GetCase(r.Config.GetString("repo"), "master", labName)
	if code != 200 {
		r.Log.MainLog.Error().Msg("Error getting lab info")
		logger.Console("Error getting lab info").Error()
		os.Exit(1)
	}

	specVM := fmt.Sprintf(r.ConfigPath+"/vagrant/%s", cases.Spec)

	client, err := vagrant.NewVagrantClient(specVM)
	if err != nil {
		r.Log.MainLog.Error().Msg(err.Error())
		os.Exit(1)
	}

	destroyCommand := client.Destroy()

	ok := destroyCommand.Run()
	if ok != nil {
		r.Log.MainLog.Error().Msg(ok.Error())
		logger.Console("Vagrant destroy failed").Error()
		os.Exit(1)
	}

	if destroyCommand.Error != nil {
		r.Log.MainLog.Error().Msg(destroyCommand.ErrorResponse.Error.Error())
		logger.Console("Vagrant destroy failed").Error()
		os.Exit(1)
	}

	logger.Console("Vagrant destroy success").Success()
}

func (r Runner) serverCreate(labName string, verbose bool) {
	logger.Console("Creating server " + labName).Start()

	cases, code := api.GetCase(r.Config.GetString("repo"), "master", labName)
	if code != 200 {
		r.Log.MainLog.Error().Msg("Error getting lab info")
		logger.Console("Error getting lab info").Error()
		os.Exit(1)
	}

	specVM := fmt.Sprintf(r.ConfigPath+"/vagrant/%s", cases.Spec)

	client, err := vagrant.NewVagrantClient(specVM)
	if err != nil {
		r.Log.MainLog.Error().Msg(err.Error())
		os.Exit(1)
	}

	// TODO: bring verbose flag from config
	// TODO: get verbose to log file
	vagrantUp := client.Up()
	vagrantUp.Verbose = verbose
	if ok := vagrantUp.Run(); ok != nil {
		r.Log.MainLog.Error().Msg(ok.Error())
		os.Exit(1)
	}

	if vagrantUp.Error != nil {
		r.Log.MainLog.Error().Msg(vagrantUp.ErrorResponse.Error.Error())
		os.Exit(1)
	}

	vmStatus, err := getVMStatus(specVM)
	if err != nil {
		r.Log.MainLog.Error().Msg(err.Error())
		os.Exit(1)
	}
	for VMName, status := range vmStatus {
		fmt.Println("    " + VMName + ": " + status)
	}

	logger.Console("Vagrant up success").Success()
}

func (r Runner) serverSSH(labName string, vmName string, sshCmd bool) {
	logger.Console("SSH server " + labName).Start()

	cases, code := api.GetCase(r.Config.GetString("repo"), "master", labName)
	if code != 200 {
		r.Log.MainLog.Error().Msg("Error getting lab info")
		logger.Console("Error getting lab info").Error()
		os.Exit(1)
	}

	specVM := fmt.Sprintf(r.ConfigPath+"/vagrant/%s", cases.Spec)

	logger.Console("Checking status of server").Info()
	if status, err := getVMStatus(specVM); err != nil {
		r.Log.MainLog.Error().Msg(err.Error())
		os.Exit(1)
	} else {
		if status[vmName] != "running" {
			logger.Console("Server is not running").Error()
			os.Exit(1)
		}
	}

	sshConfig, err := getVMSSHConfig(specVM)
	if err != nil {
		r.Log.MainLog.Error().Msg(err.Error())
		os.Exit(1)
	}

	if sshCmd {
		fmt.Println("vagrant ssh " + vmName)
		fmt.Println("or use ssh instead")
		fmt.Printf("ssh -i %s -p %d %s@%s",
			sshConfig[vmName].IdentityFile,
			sshConfig[vmName].Port,
			sshConfig[vmName].User,
			sshConfig[vmName].HostName)
		logger.Console("SSH command").Success()
		return
	}

	logger.Console("philo ssh is EXPERIMENTAL use \"--command\" instead to show ssh command").Warn()

	config := &ssh.ClientConfig{
		User: sshConfig[vmName].User,
		Auth: []ssh.AuthMethod{
			remote.PublicKeyFile(sshConfig[vmName].IdentityFile),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", sshConfig[vmName].HostName, sshConfig[vmName].Port), config)
	if err != nil {
		fmt.Println("Failed to dial:", err)
		return
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		fmt.Println("Failed to create session:", err)
		return
	}
	defer session.Close()

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	input, err := session.StdinPipe()
	if err != nil {
		fmt.Println("Failed to create input pipe:", err)
		return
	}

	go func() {
		reader := bufio.NewReader(os.Stdin)
		input.Write([]byte("echo 'Welcome to Philo'\n"))
		for {
			fmt.Print(vmName, "~> ")
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Println("Failed to read from stdin:", err)
				return
			}
			input.Write([]byte(line))
		}
	}()

	err = session.Shell()
	if err != nil {
		fmt.Println("Failed to start shell:", err)
		return
	}

	err = session.Wait()
	if err != nil {
		fmt.Println("Failed to wait for session:", err)
		return
	}

	logger.Console("SSH server success").Success()
}

func serverCommand() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "server",
		Short: "manage your local server",
		Long: `philo use vagrant to manage your local server, command replacement for serveral vagrant command
vagrant up, vagrant destroy, vagrant halt, vagrant status`,
	}

	serverSSH := &cobra.Command{
		Use:   "ssh [lab name] [vm name]",
		Short: "use philo to ssh into your server [EXPERIMENTAL]",
		Long: `connect to your server using philo server ssh and return remote shell, 
for stability use "philo server ssh --command" to show ssh command`,
		Args: cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			sshCmd, err := cmd.Flags().GetBool("command")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			runner.serverSSH(args[0], args[1], sshCmd)
		},
	}

	serverSSH.Flags().BoolP("command", "c", false, "show ssh command")

	serverUp := &cobra.Command{
		Use:   "up [LAB NAME]",
		Short: "bringing up your server",
		Long:  `philo server up will bring up your server, that is same as vagrant up`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			verbose, err := cmd.Flags().GetBool("verbose")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			runner.serverCreate(args[0], verbose)
		},
	}
	// create flag for verbose
	serverUp.Flags().BoolP("verbose", "v", false, "show vagrant output")

	serverShow := &cobra.Command{
		Use:   "show [LAB NAME]",
		Short: "show your server status",
		Long:  `philo server show will show your server status, that is same as vagrant status`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runner.serverShow(args[0])
		},
	}

	serverHalt := &cobra.Command{
		Use:   "halt [LAB NAME]",
		Short: "halt/shutdown your server",
		Long: `philo server halt will halt/shutdown your server, that is same as vagrant halt

		make sure shutdown server before shutdown your computer to avoid data corruption`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runner.serverHalt(args[0])
		},
	}

	serverDestroy := &cobra.Command{
		Use:   "destroy [LAB NAME]",
		Short: "destroy your server",
		Long:  `philo server destroy will destroy your server, that is same as vagrant destroy`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runner.serverDestroy(args[0])
		},
	}

	cmd.AddCommand(serverSSH, serverUp, serverShow, serverHalt, serverDestroy)

	return cmd
}
