package utils

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var templateTs = &cobra.Command{
	Use:   "template-ts",
	Short: "Generate node ts project using Typescript",
	Run:   generateNodeTsTemplate,
}

func init() {
	rootCmd.AddCommand(templateTs)
}
func generateNodeTsTemplate(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("Error: Please provide a folder name for the template.")
		return
	}
	repoUrl := "https://github.com/faanrm/nodeTs-template"
	destDir := args[0] //using to get the foldername to command line
	CreateDirectory(destDir)
	fmt.Println("Waiting for generating template .....")

	//clone repo first
	cmdClone := exec.Command("git", "clone", repoUrl)
	cmdClone.Stdout = os.Stdout
	cmdClone.Stderr = os.Stderr
	err := cmdClone.Run()
	if err != nil {
		fmt.Println("Error when generating template:", err)
		return
	}
	//move all repo to the desired folder
	MoveFiles(cmd, destDir, "nodeTs-template")
	fmt.Println("Template generated successfully")
}
