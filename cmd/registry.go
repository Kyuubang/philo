package cmd

import (
	"github.com/Kyuubang/philo/logger"
	"github.com/spf13/viper"
)

// Runner struct for organizing all runner properties
type Runner struct {
	// ConfigPath is the path to the config file directory
	ConfigPath string
	// Config is the viper instance $HOME/.philo/config.json
	Config *viper.Viper
	// Log is the logger instance
	Log Log
}

type Log struct {
	// MainLog is the logger instance for the main logger /var/log/philo/philo.log
	MainLog *logger.Logger
	// LabLog is the logger instance for the lab logger /var/log/philo/labs/<name>.log
	LabLog *logger.Logger
}

var runner = Runner{
	Config:     initConfig(),
	ConfigPath: createConfDir(),
	Log: Log{
		MainLog: logger.MainLog(),
	},
}

func init() {

	// register all subcommands
	philoCmd.AddCommand(initCommand())
	philoCmd.AddCommand(labCommand())
	philoCmd.AddCommand(serverCommand())
	philoCmd.AddCommand(scoreCommand())
	philoCmd.AddCommand(configCommand())
	philoCmd.AddCommand(loginCommand())
}
