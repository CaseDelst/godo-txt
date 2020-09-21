package cmd

import (
	"github.com/spf13/viper"
)

type config struct {
	// JSONMode is the initial JSON mode to use
	FilePath string // `json:"filePath"`
}

var configFile map[string]config
var filePath string
var configFileName string
var currentConfig *string
var configViper *viper.Viper

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	filePath = "C:\\Users\\admin\\Google Drive\\todo.txt"

	/* if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		fmt.Printf("homedir:%s", home)
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
	fmt.Println("Using config file:", viper.ConfigFileUsed()) */
}

/* func reloadConfigFile() {
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
} */

/* func addConfigCommands(rootCmd *cobra.Command) {
	root := &cobra.Command{
		Use:     "configs",
		Short:   "update configurations",
		Aliases: []string{"cfg"},
		RunE: func(cmd *cobra.Command, args []string) error {
			config, err := listConfigs()
			if err != nil {
				return err
			}
			fmt.Println("--Current Config(s):--")
			for i, c := range config {
				fmt.Printf("%v:\nFilePath: %s\n", i, c.FilePath)
			}
			return nil
		},
	}

	rootCmd.AddCommand(root)
} */

/* func setConfig(name string, config config) {
	currentConfig = &name
}

func listConfigs() (map[string]config, error) {
	return configFile, nil
}

func listConfigKeys() ([]string, error) {
	configs, err := listConfigs()
	if err != nil {
		return nil, err
	}
	keys := make([]string, 0, len(configs))
	for c := range configs {
		keys = append(keys, c)
	}

	return keys, nil
}
*/
