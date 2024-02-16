package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var template = &cobra.Command{
	Use:   "template",
	Short: "Generate nodeJS template",
	Long:  `Generate node template`,
	Run:   generateTemplate,
}

func init() {
	rootCmd.AddCommand(template)
	// Modify the template command to accept arguments for the folder name and authentication flag
	template.Flags().StringP("directory", "D", "./myApp", "Output directory for the project")
	template.Flags().StringP("database", "d", "mongo", "Choose database type: mongo or sql")
	template.Flags().BoolP("auth", "a", false, "Include authentication")

}

func generateTemplate(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("Error: Please provide a folder name for the template.")
		return
	}

	// Check if the database flag is provided
	dbFlag := cmd.Flag("database")
	if !dbFlag.Changed {
		// If database flag is not provided, print an error and show the command to specify database
		fmt.Println("Error: Please specify the database type using the --database flag.")
		fmt.Println("Example: gno template --database <mongo/sql> <folder_name>")
		return
	}

	// Continue with template generation

	repoURLMongo := "https://github.com/Faanilo/API-EXPRESS.git"
	repoURLSQL := "https://github.com/faanrm/NodeJs-Sequelize-Starter.git"
	destDir := args[0]

	// Get the selected database type from command line flags
	dbType, _ := cmd.Flags().GetString("database")
	var repoURL string

	// Choose the appropriate repoURL based on the selected database type
	switch dbType {
	case "mongo":
		repoURL = repoURLMongo
	case "sql":
		repoURL = repoURLSQL
	default:
		fmt.Println("Error: Invalid database type. Please choose 'mongo' or 'sql'.")
		return
	}

	// Check if the auth flag is provided
	authFlag, _ := cmd.Flags().GetBool("auth")
	if authFlag {
		// If auth flag is provided, clone the repository and then switch to the "auth" branch
		err := cloneAndCheckout(repoURL, "auth", destDir)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	} else {
		// If auth flag is not provided, clone the repository without switching branches
		err := cloneAndCheckout(repoURL, "main", destDir) // Assuming main is the main branch name, change it if necessary
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	// Remove the .git directory
	err := os.RemoveAll(filepath.Join(destDir, ".git"))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\nTemplate generated successfully")
}

func cloneAndCheckout(repoURL, branch, destDir string) error {
	// Create the destination directory if it doesn't exist
	CreateDirectoryTemp(destDir)

	// Clone the repository
	cmdClone := exec.Command("git", "clone", "-b", branch, "--single-branch", repoURL, destDir)
	cmdClone.Stdout = os.Stdout
	cmdClone.Stderr = os.Stderr
	err := cmdClone.Run()
	if err != nil {
		return err
	}

	// Change to the specified branch
	cmdCheckout := exec.Command("git", "-C", destDir, "checkout", branch)
	cmdCheckout.Stdout = os.Stdout
	cmdCheckout.Stderr = os.Stderr
	err = cmdCheckout.Run()
	if err != nil {
		return err
	}

	return nil
}

/*func CreateDirectoryTemp(destDir string) {
	err := os.MkdirAll(destDir, os.ModePerm)
	if err != nil {
		fmt.Println("Error:", err)
	}
}*/
