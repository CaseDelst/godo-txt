package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

type config struct {
	// JSONMode is the initial JSON mode to use
	FilePath string // `json:"filePath"`
}

var configFile map[string]config
var configFileName string
var configViper *viper.Viper

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			er(err)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".godotxt")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	fmt.Println("Using config file:", viper.ConfigFileUsed())
}

func reloadConfigFile() {
	if err := configViper.ReadInConfig(); err != nil {
		configFileExpected := configFileName != ""
		_, errIsNotFound := err.(viper.ConfigFileNotFoundError)
		if !configFileExpected && errIsNotFound {
			return
		}
		fmt.Fprintf(os.Stderr, "Error loading config file: %s\n", err)
		os.Exit(1)
	}

	configFile = nil
	if err := configViper.Unmarshal(&configFile); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse config file: %s", err)
		os.Exit(1)
	}
}
