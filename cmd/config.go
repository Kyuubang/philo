package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"path"
)

// ConfigFile is default structure config.json file
type ConfigFile struct {
	Repo  string `json:"repo,omitempty"`
	Auths Auths  `json:"auths,omitempty"`
}
type GithubCom struct {
	Username string `json:"username,omitempty"`
	Token    string `json:"token,omitempty"`
}
type Auths struct {
	GithubCom GithubCom `json:"github.com,omitempty"`
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
	var dataconfig ConfigFile
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
