package utils

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var blue = "\033[34m"
var reset = "\033[0m"

var (
	versionFlag bool
	version     = "0.1.2"
)

var rootCmd = &cobra.Command{
	Use: "gno",
	Short: blue + `
_____/\\\\\\\\\\\\__/\\\\\_____/\\\_______/\\\\\______        
___/\\\//////////__\/\\\\\\___\/\\\_____/\\\///\\\____       
__/\\\_____________\/\\\/\\\__\/\\\___/\\\/__\///\\\__      
_\/\\\____/\\\\\\\_\/\\\//\\\_\/\\\__/\\\______\//\\\_     
_\/\\\___\/////\\\_\/\\\\//\\\\/\\\_\/\\\_______\/\\\_    
_\/\\\_______\/\\\_\/\\\_\//\\\/\\\_\//\\\______/\\\__   
_\/\\\_______\/\\\_\/\\\__\//\\\\\\__\///\\\__/\\\____  
_\//\\\\\\\\\\\\/__\/\\\___\//\\\\\____\///\\\\\/_____ 
__\////////////____\///_____\/////_______\/////_______
` + reset,
	Run: func(cmd *cobra.Command, args []string) {
		if versionFlag {
			fmt.Println("Version:", version)
			return
		}
		err := cmd.Help()
		if err != nil {
			fmt.Println("Error displaying help:", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&versionFlag, "version", "v", false, "See version installed")
}
