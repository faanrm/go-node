package utils

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gno",
	Short: "Gno Code Generator",
	Long: `Welcome to GNO: Your Node Project Generator!
	GNO is a powerful and easy-to-use command line tool designed to kickstart your Node.js development. 
	With just a few flags, you can generate a structured Node project, 
	saving you time and setting you up with best practices.`,
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
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
