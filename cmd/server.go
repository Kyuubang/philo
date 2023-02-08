package cmd

import (
	"fmt"
	"github.com/Kyuubang/philo/logger"
	"github.com/bmatcuk/go-vagrant"
	"github.com/spf13/cobra"
	"os"
)

// TODO: Create replacement command for vagrant up, ssh, destroy, etc

func (r Runner) serverSSH(labName string) {

}

func (r Runner) serverShow(labName string) {
	logger.Console("Showing server").Start()

	client, err := vagrant.NewVagrantClient("/home/bayhaqi/Repository/Github/philo-sample-case/linux/linux-001-1/vagrant")
	if err != nil {
		r.MainLog.Error().Msg(err.Error())
		os.Exit(1)
	}

	statusCommand := client.Status()

	ok := statusCommand.Run()
	if ok != nil {
		r.MainLog.Error().Msg(ok.Error())
		os.Exit(1)
	}

	if statusCommand.Error != nil {
		r.MainLog.Error().Msg(statusCommand.ErrorResponse.Error.Error())
		os.Exit(1)
	}

	for machine, status := range statusCommand.StatusResponse.Status {
		fmt.Printf("    %s: %s\n", machine, status)
	}
	logger.Console("Vagrant status success").Success()
}

func (r Runner) serverHalt(labName string) {
	logger.Console("Halting server").Start()

	client, err := vagrant.NewVagrantClient("/home/bayhaqi/Repository/Github/philo-sample-case/linux/linux-001-1/vagrant")
	if err != nil {
		r.MainLog.Error().Msg(err.Error())
		os.Exit(1)
	}

	haltCommand := client.Halt()

	ok := haltCommand.Run()
	if ok != nil {
		r.MainLog.Error().Msg(ok.Error())
		logger.Console("Vagrant halt failed").Error()
		os.Exit(1)
	}

	if haltCommand.Error != nil {
		r.MainLog.Error().Msg(haltCommand.ErrorResponse.Error.Error())
		os.Exit(1)
	}

	logger.Console("Vagrant halt success").Success()
}

func (r Runner) serverDestroy(labName string) {
	logger.Console("Destroying server " + labName).Start()

	client, err := vagrant.NewVagrantClient("/home/bayhaqi/Repository/Github/philo-sample-case/linux/linux-001-1/vagrant")
	if err != nil {
		r.MainLog.Error().Msg(err.Error())
		os.Exit(1)
	}

	destroyCommand := client.Destroy()

	ok := destroyCommand.Run()
	if ok != nil {
		r.MainLog.Error().Msg(ok.Error())
		os.Exit(1)
	}

	if destroyCommand.Error != nil {
		r.MainLog.Error().Msg(destroyCommand.ErrorResponse.Error.Error())
		os.Exit(1)
	}

	logger.Console("Vagrant destroy success").Success()
}

func (r Runner) serverCreate(labName string) {
	logger.Console("Creating server " + labName).Start()
	client, err := vagrant.NewVagrantClient("/home/bayhaqi/Repository/Github/philo-sample-case/linux/linux-001-1/vagrant")
	if err != nil {
		r.MainLog.Error().Msg(err.Error())
		os.Exit(1)
	}

	vagrantUp := client.Up()
	vagrantUp.Verbose = true
	ok := vagrantUp.Run()
	if ok != nil {
		r.MainLog.Error().Msg(ok.Error())
		os.Exit(1)
	}

	if vagrantUp.Error != nil {
		r.MainLog.Error().Msg(vagrantUp.ErrorResponse.Error.Error())
		os.Exit(1)
	}

	logger.Console("Vagrant up success").Success()

	response := vagrantUp.UpResponse
	for index, _ := range response.VMInfo {
		fmt.Println("philo server ssh", index)
	}
}

func serverCommand() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "server",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	}

	cmd.AddCommand(
		&cobra.Command{
			Use:   "create",
			Short: "A brief description of your command",
			Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
			Args: cobra.MinimumNArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				runner.serverCreate(args[0])
			},
		},
		&cobra.Command{
			Use:   "show",
			Short: "A brief description of your command",
			Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
			Args: cobra.MinimumNArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				runner.serverShow(args[0])
			},
		},
		&cobra.Command{
			Use:   "halt",
			Short: "A brief description of your command",
			Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
			Args: cobra.MinimumNArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				runner.serverHalt(args[0])
			},
		},
		&cobra.Command{
			Use:   "destroy",
			Short: "A brief description of your command",
			Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
			Args: cobra.MinimumNArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				runner.serverDestroy(args[0])
			},
		})

	return
}
