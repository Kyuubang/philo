package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"path"
)

// ConfigFile is default structure config.json file
type ConfigFile struct {
	Repo  string            `json:"repo"`
	Auths map[string]string `json:"auths"`
}
type GithubCom struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

// create config directory on $HOME/.philo
func createConfDir() string {
	home, _ := os.UserHomeDir()
	configPath := path.Join(home, ".philo")
	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		err = os.Mkdir(configPath, 0700)
		if err != nil {
			fmt.Println(err)
		}
	}
	return configPath
}

// create default config file -> $HOME/.philo/config.json
func createDefConf() {
	var dataconfig = ConfigFile{
		Repo: "github.com/example/example",
	}
	file, _ := json.MarshalIndent(dataconfig, "", "\t")
	pathconf := createConfDir()
	err := ioutil.WriteFile(pathconf+"/config.json", file, 0600)
	if err != nil {
		fmt.Println("unable to write default config.json")
	}
}

// initConfig for main config (.philo/config.json) and hosts config (.philo/hosts.yaml)
func initConfig() *viper.Viper {
	var mainConfig = viper.New()
	mainConfig.SetConfigName("config")
	mainConfig.SetConfigType("json")
	mainConfig.AddConfigPath("$HOME/.philo")
	err := mainConfig.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			createDefConf()
			fmt.Println("config file not found! created one.")
			fmt.Println("config file not found! created one.", "warn")
		} else {
			fmt.Println("failed when read $HOME/.philo/config.json file")
			fmt.Println("failed when read $HOME/.philo/config.json file", "warn")
			os.Exit(1)
		}
	}
	return mainConfig
}

// create runner command to modify config file
func (r Runner) setConfig(key string, value string) {
	r.Config.Set(key, value)
	if err := r.Config.WriteConfig(); err != nil {
		fmt.Println(err)
	}
}

func configCommand() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "config",
		Short: "config is a command to manage config file",
		Long:  `config is a command to manage config file`,
	}

	var setCmd = &cobra.Command{
		Use:   "set [KEY] [VALUE]",
		Short: "set config file by key and value",
		Long:  `set is a command to set config file`,
		Run: func(cmd *cobra.Command, args []string) {
			runner.setConfig(args[0], args[1])
		},
	}

	cmd.AddCommand(setCmd)

	return
}
