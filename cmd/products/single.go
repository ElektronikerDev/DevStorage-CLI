package products

import (
	"github.com/spf13/cobra"
)

func SingleProduct(apiKeyString string) *cobra.Command {
	apiKey = apiKeyString
	return productCmd
}

var productCmd = &cobra.Command{
	Use:                   "product",
	Short:                 "Get Product Information",
	Long:                  `Get information about a product.`,
	DisableAutoGenTag:     true,
	DisableSuggestions:    true,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		runProduct(cmd, args)
	},
}

func runProduct(cmd *cobra.Command, args []string) error {
	//TODO: get Product Information
	return nil
}
