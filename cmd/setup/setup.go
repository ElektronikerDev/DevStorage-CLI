package setup

import (
	"bufio"
	"encoding/json"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"time"
)

func Setup() *cobra.Command {
	return setupCmd
}

var setupCmd = &cobra.Command{
	Use:                   "setup",
	Short:                 "DevStorage CLI Setup",
	Long:                  `The setup command is used to set up the CLI. In the setup the required CLI-API key is requested, with which later a simple control on the server and other products of the customer can be granted.`,
	DisableAutoGenTag:     true,
	DisableSuggestions:    true,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		runSetup(cmd, args)
	},
}

type CLIUserObject struct {
	Firstname string
	Lastname  string
}
type ApiCheckCall struct {
	Successfully bool
	User         CLIUserObject
}

func runSetup(cmd *cobra.Command, args []string) error {
	log.Println("CLI Setup starting..")
	// Get the linux user. Needed for the home directory
	linuxUser, _ := user.Current()
	keyFile := linuxUser.HomeDir + "\\dvstrg_cli.key"
	if _, err := ioutil.ReadFile(keyFile); err == nil {
		// file already exist
		log.Println("API Key already exist.")
		log.Fatalln("If you want to save a new key, delete the file: " + keyFile)
	}

	// This check is for the internet connection and the availablity of the CLI Rest Endpoint
	res, err := http.Get("https://rest.devstorage.eucli")

	// Throw Fatal Log if the connection fail
	if err != nil {
		log.Fatalln("Setup cannot be executed because the application cannot connect to the DevStorage API interface.")
	}
	res.Body.Close()

	// -> Successfully connected

	log.Println("Successfully connected to the API interface of DevStorage\n\n")
	log.Println("Create your CLI-API key here: https://hosting.devstorage.eu/settings/cli")
	log.Print("Please enter your CLI-API key here: ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	var cliKey = input.Text()

	// Client-Side Key validation
	if len(cliKey) != 16 {
		log.Printf("Invalid CLI Key! Do you need help with this? Contact our lovely support.")
		return nil
	}

	// Check API-Key is valid via Rest API
	apiCallCheck := ApiCheckCall{}
	getJson("https://rest.devstorage.eucli/key/"+cliKey, &apiCallCheck)

	if !apiCallCheck.Successfully {
		log.Fatalln("It seems that the API is invalid or you are not authorized to use the CLI on this server. " +
			"\nIf you have any questions, please feel free to contact our support.")
		return nil
	}

	log.Printf("Hello " + apiCallCheck.User.Firstname + " your API Key is correct.")

	// Write KeyFile to Homefolder (keyFile)
	d1 := []byte(cliKey)
	err = ioutil.WriteFile(keyFile, d1, 0644)
	if err != nil {
		log.Fatalln("The API key could not be saved. Try again or check if the software can write to the $Home directory. (" + keyFile + ")")
		return err
	}
	log.Printf("Your API key has been saved and you can use the CLI now :)")

	// Setup finished
	return nil
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
