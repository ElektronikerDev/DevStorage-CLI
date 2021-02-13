package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os/user"

	"./setup"
	"./version"
)

var RESTAPI = "https://rest.devstorage.eu"

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:     "dscloud",
		Short:   "DevStorage Cloud CLI",
		Long:    "A command-line system for DevStorage.eu",
		Version: "1.0PA",

		TraverseChildren:      true,
		SilenceUsage:          true,
		SilenceErrors:         true,
		DisableFlagsInUseLine: true,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {

	rootCmd.PersistentFlags().StringP("author", "a", "Thomas Michaelis", "Software author")

	// Register Commands
	rootCmd.AddCommand(setup.Setup())
	rootCmd.AddCommand(version.Version())

	// Get Os User who start the software
	osUser, _ := user.Current()
	if _, err := ioutil.ReadFile(osUser.HomeDir + "\\dvstrg_cli.key"); err != nil {
		color.Red("[!] Please run the setup to be able to use all commands! [dscloud setup] [!]")
		return
	}

	//register all commands (only if cli key file exist)
}
