package utils

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var genNode = &cobra.Command{
	Use:   "node-gen",
	Short: "Generate nodejs project",
	Long: `Generate nodejs project
	Example usage:
	gno node-gen -D ./projects/myapp -l "lib1 lib2" -d "lib3 lib4"`,
	Run: generateNodeJS,
}

var template = &cobra.Command{
	Use:   "template",
	Short: "Generate CRUD node template and specify the folder name",
	Long:  `Generate CRUD node template`,
	Run:   generateTemplate,
}

func init() {
	rootCmd.AddCommand(genNode)
	rootCmd.AddCommand(template)
	genNode.Flags().StringP("directory", "D", "./myApp", "Output directory for the project")
	genNode.Flags().BoolP("yes", "y", false, "Generate default NodeJs package.json file")
	genNode.Flags().StringP("libs", "l", " ", "List of Node.js libraries to install")
	genNode.Flags().StringP("dev-libs", "d", " ", "List of Node.js libraries to install")
	// Modify the template command to accept an argument for the folder name
	template.Flags().StringP("directory", "D", "./myApp", "Output directory for the project")
}

func generateNodeJS(cmd *cobra.Command, args []string) {
	dir, _ := cmd.Flags().GetString("directory")
	createDirectory(dir)
	checkNPMInstallation()
	changeDirectory(dir)
	initNodeProject(cmd)
	installLibraries(cmd, "libs")
	installLibraries(cmd, "dev-libs")
}

func generateTemplate(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("Error: Please provide a folder name for the template.")
		return
	}

	repoURL := "https://github.com/Faanilo/API-EXPRESS.git"
	destDir := args[0] // Get the folder name from the command line arguments

	// Create the destination directory if it doesn't exist
	createDirectoryTemp(destDir)
	fmt.Println("Generating template ...")

	// Clone the repository
	cmdClone := exec.Command("git", "clone", repoURL)
	cmdClone.Stdout = os.Stdout
	cmdClone.Stderr = os.Stderr
	err := cmdClone.Run()

	if err != nil {
		fmt.Println("Error when generating template:", err)
		return
	}

	// Move all files from the cloned repository to the destination directory
	moveFiles(cmd, destDir, "API-EXPRESS")

	fmt.Println("Template generated successfully ")
}
