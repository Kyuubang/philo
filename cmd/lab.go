package cmd

import (
	"fmt"
	"github.com/Kyuubang/philo/cmd/ui"
	"github.com/Kyuubang/philo/internal/api"
	"github.com/Kyuubang/philo/logger"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"os"
	"path"
	"strings"
)

func (r Runner) labList(course string) {
	var repo = strings.Split(r.Config.GetString("repo"), "/")
	if len(repo) < 3 {
		r.Log.MainLog.Error().Msg("invalid repo config")
		logger.Console("invalid repo config").Error()
		os.Exit(1)
	}
	var items, code = api.GetListLab(path.Join(repo[1], repo[2]), course)

	switch code {
	case 404:
		r.Log.MainLog.Error().Msg("Course not found")
		logger.Console("Course not found").Error()
		os.Exit(1)
	case 403:
		r.Log.MainLog.Error().Msg("Forbidden")
		logger.Console("Forbidden").Error()
		os.Exit(1)
	case 200:
		//fmt.Println("[+] Available lab for " + course)
		logger.Console("Available lab for " + course).Start()

		for index, item := range items {
			fmt.Printf("    %d. %s\n", index+1, item)
		}
	}
}

func (r Runner) labView(lab string) {
	var repo = strings.Split(r.Config.GetString("repo"), "/")
	var instruction, code = api.GetReadme(path.Join(repo[1], repo[2]), "master", lab)

	switch code {
	case 404:
		logger.Console("Lab not found").Error()
		os.Exit(1)
	case 403:
		logger.Console("Forbidden").Warn()
		os.Exit(1)
	}

	model, err := ui.MDReader(string(instruction))
	if err != nil {
		fmt.Println("Could not initialize Bubble Tea model:", err)
		os.Exit(1)
	}

	if _, err := tea.NewProgram(model).Run(); err != nil {
		fmt.Println("Bummer, there's been an error:", err)
		os.Exit(1)
	}
}

// labCheck will check your work based on the lab you are doing
func (r Runner) labCheck(lab string) {

}

func labCommand() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "lab",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	}

	labList := &cobra.Command{
		Use:   "list",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runner.labList(args[0])
		},
	}

	labView := &cobra.Command{
		Use:   "view",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runner.labView(args[0])
		},
	}

	cmd.AddCommand(labList, labView)

	return cmd
}
