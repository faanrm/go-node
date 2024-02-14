package utils

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "gno",
	Short: `
_____/\\\\\\\\\\\\__/\\\\\_____/\\\_______/\\\\\______        
 ___/\\\//////////__\/\\\\\\___\/\\\_____/\\\///\\\____       
  __/\\\_____________\/\\\/\\\__\/\\\___/\\\/__\///\\\__      
   _\/\\\____/\\\\\\\_\/\\\//\\\_\/\\\__/\\\______\//\\\_     
    _\/\\\___\/////\\\_\/\\\\//\\\\/\\\_\/\\\_______\/\\\_    
     _\/\\\_______\/\\\_\/\\\_\//\\\/\\\_\//\\\______/\\\__   
      _\/\\\_______\/\\\_\/\\\__\//\\\\\\__\///\\\__/\\\____  
       _\//\\\\\\\\\\\\/__\/\\\___\//\\\\\____\///\\\\\/_____ 
        __\////////////____\///_____\/////_______\/////_______
		`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags()
}
