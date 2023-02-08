package cmd

import (
	"github.com/Kyuubang/philo/logger"
	"github.com/spf13/viper"
)

// Runner struct for organizing all runner properties
type Runner struct {
	ConfigPath string
	Config     *viper.Viper
	MainLog    *logger.Logger
	LabLog     *logger.Logger
}

var runner = Runner{
	Config:     initConfig(),
	ConfigPath: createConfDir(),
	MainLog:    logger.MainLog(),
}

func init() {
	// register all subcommands
	philoCmd.AddCommand(initCommand())
	philoCmd.AddCommand(labCommand())
	philoCmd.AddCommand(serverCommand())
	philoCmd.AddCommand(scoreCommand())

}
