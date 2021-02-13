package products

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"strconv"
)

func Products(apiKeyString string) *cobra.Command {
	apiKey = apiKeyString
	return productsCmd
}

var productsCmd = &cobra.Command{
	Use:                   "products",
	Short:                 "List all your products",
	Long:                  `It will list all the products you currently own.`,
	DisableAutoGenTag:     true,
	DisableSuggestions:    true,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		runProductList(cmd, args)
	},
}

func runProductList(cmd *cobra.Command, args []string) error {

	cliProducts := CLIProducts{}
	getJson("http://localhost:9999/cli/products/"+apiKey, &cliProducts)
	color.HiBlue("===================[ Products ]===================")
	for i, product := range cliProducts.Products {
		fmt.Println("[" + strconv.Itoa(i) + "] " + product.Name + "-" + strconv.Itoa(product.Id) + " (" + product.ExpirationDate + ")")
	}
	color.HiBlue("==================================================")
	return nil
}
