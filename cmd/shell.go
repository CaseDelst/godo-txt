package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/spf13/cobra"
)

func runShell(cmd *cobra.Command, args []string) {
	addShellCommands(rootCmd)

	fmt.Printf("Path")
	p := prompt.New(
		Executor,
		completer,
		prompt.OptionTitle("godotxt: Go todo.txt CLI"),
		prompt.OptionLivePrefix(createPrefix),
		prompt.OptionPrefixTextColor(prompt.Cyan),
		prompt.OptionInputTextColor(prompt.DarkGreen),
		prompt.OptionDescriptionBGColor(prompt.LightGray),
		prompt.OptionSuggestionBGColor(prompt.DarkGray),
		prompt.OptionScrollbarThumbColor(prompt.Cyan),
		prompt.OptionScrollbarBGColor(prompt.Black),
		prompt.OptionSelectedDescriptionBGColor(prompt.Cyan),
		prompt.OptionSelectedSuggestionBGColor(prompt.Cyan),
		prompt.OptionSelectedSuggestionTextColor(prompt.LightGray),
	)
	p.Run()
}

func addShellCommands(rootCmd *cobra.Command) {
	rootCmd.AddCommand(&cobra.Command{
		Use:     "pwd",
		Aliases: []string{"status", "current"},
		Short:   "show current context (api key, project, environment)",
		Run:     printCurrentSettings,
	})

	
}

func completer(d prompt.Document) []prompt.Suggest {
	return prompt.FilterFuzzy(createSuggestions(rootCmd, d), d.GetWordBeforeCursor(), true)
}

func createSuggestions(c *cobra.Command, p prompt.Document) []prompt.Suggest {
	inputArgs := strings.Split(p.CurrentLine(), " ")
	var suggestions []prompt.Suggest

	switch len(inputArgs) {
	case 1:
		for _, child := range c.Commands() {
			suggestions = append(suggestions, prompt.Suggest{
				Text:        child.Use,
				Description: child.Short,
			})
		}
		suggestions = addUtilitySuggestions(suggestions)
	case 2:
		for _, child := range c.Commands() {
			if contains(inputArgs, child.Use) {
				for _, grandchild := range child.Commands() {
					suggestions = append(suggestions, prompt.Suggest{
						Text:        grandchild.Use,
						Description: grandchild.Short,
					})
				}
			}
		}
	}

	return suggestions
}

func printCurrentSettings(cmd *cobra.Command, args []string) {
	fmt.Println("Current Config: " + noneIfNil(currentConfig))
	fmt.Println("Current Path: " + configFile[*currentConfig].FilePath)
}

func addUtilitySuggestions(suggestions []prompt.Suggest) []prompt.Suggest {
	return append(suggestions,
		prompt.Suggest{
			Text:        "help",
			Description: "more information on any command",
		},
		prompt.Suggest{
			Text:        "quit",
			Description: "exit the shell",
		},
		prompt.Suggest{
			Text:        "exit",
			Description: "exit the shell",
		},
	)
}

func Executor(s string) {
	if s == "" {
		return
	} else if s == "exit" || s == "quit" {
		fmt.Println("Exiting!")
		os.Exit(0)
		return
	}

	rootCmd.SetArgs(strings.Fields(s))
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func createPrefix() (string, bool) {
	if currentConfig != nil {
		return fmt.Sprintf("<%s> ", *currentConfig), true
	}

	return "<<>> ", true
}

func noneIfNil(s *string) string {
	if s == nil {
		return "<none>"
	}
	return *s
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
