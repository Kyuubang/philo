package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var philoCmd = &cobra.Command{
	Use:   "philo",
	Short: "Test case executor for your labs",
	Long: `Philo is a test case executor for your labs. It is designed to be used with
the philo server, which is used to manage your labs and their test cases.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := philoCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
