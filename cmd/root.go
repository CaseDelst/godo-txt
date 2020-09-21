package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile string

	rootCmd = &cobra.Command{
		Use:              "godo",
		Short:            "A CLI for the todo.txt format",
		// PersistentPreRun: preRunCmd,
	}

	shellCmd = &cobra.Command{
		Use:    "shell",
		Short:  "start an interactive shell",
		// PreRun: preRunCmd,
		Run: func(cmd *cobra.Command, args []string) {
			runShell(cmd, args)
		},
	}
)

// Execute executes the root command.
func Execute() {
	cmd := rootCmd
	rootCmdWithShell := *rootCmd
	rootCmdWithShell.AddCommand(shellCmd)
	foundCmd, _, err := rootCmdWithShell.Find(os.Args[1:])
	if err == nil && foundCmd.Use == "shell" {
		cmd = shellCmd
	}
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
}

/* func preRunCmd(cmd *cobra.Command, args []string) {
	if currentConfig == nil {
		configs, err := listConfigs()
		if err != nil {
			configs = nil
		}

		config := viper.GetString("config")

		// Assume we can use the single config if there is only one
		if config == "" && len(configs) == 1 {
			for name := range configs {
				config = name
				break
			}
		}

		if config != "" {
			found := false
			for name, v := range configs {
				if name == config {
					setConfig(name, v)
					found = true
					break
				}
			}
			if !found {
				fmt.Fprintf(os.Stderr, `Unable to find config "%s"`, config)
				os.Exit(1)
			}
		}

		if viper.IsSet("path") {
			if token := viper.GetString("token"); token != "" {
				cfgFile = token
			}
		}
	}
} */

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}
