package cmd

import (
	"github.com/Kyuubang/philo/logger"
	"github.com/spf13/cobra"
)

func (r Runner) scoreCheck(labName string) {
	logger.Console("Checking score").Start()

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
