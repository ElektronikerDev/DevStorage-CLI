package version

import (
	"github.com/spf13/cobra"
	"log"
)

func Version() *cobra.Command {
	return versionCmd
}

var versionCmd = &cobra.Command{
	Use:                   "version",
	Short:                 "DevStorage CLI Version",
	Long:                  `Print the version number of the DevStorage CLI`,
	DisableAutoGenTag:     true,
	DisableSuggestions:    true,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("DevStorage CLI Version 1.0PA by Thomas Michaelis for DevStorage")
	},
}
