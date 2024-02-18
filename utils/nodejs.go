package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"

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

	repoURLMongo := "https://github.com/faanrm/NodeJs-Mongo-Starter.git"
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

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		// Run git clone in background
		err := cloneAndCheckout(repoURL, destDir, authFlag)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}()

	// Print progress percentage
	const totalSteps = 10 // Adjust this based on your estimation of the steps in git clone
	for i := 0; i <= totalSteps; i++ {
		// Print progress bar
		const progressBarWidth = 50
		progress := i * progressBarWidth / totalSteps
		bar := strings.Repeat("=", progress) + strings.Repeat(" ", progressBarWidth-progress)
		fmt.Printf("\rProgress: [%s] %d%%", bar, (i*100)/totalSteps)

		os.Stdout.Sync()
		time.Sleep(500 * time.Millisecond) // Adjust sleep duration as needed
	}
	fmt.Println("\nTemplate generated successfully")

	wg.Wait()
}

func cloneAndCheckout(repoURL, destDir string, authFlag bool) error {
	// Create the destination directory if it doesn't exist
	CreateDirectoryTemp(destDir)

	// Create a buffer to capture the output
	var outputBuffer bytes.Buffer

	// Clone the repository
	cmdClone := exec.Command("git", "clone", "-b", "main", "--single-branch", "--depth", "1", repoURL, destDir)
	if authFlag {
		cmdClone = exec.Command("git", "clone", "-b", "auth", "--single-branch", "--depth", "1", repoURL, destDir)
	}

	// Redirect both stdout and stderr to the buffer
	cmdClone.Stdout = &outputBuffer
	cmdClone.Stderr = &outputBuffer

	err := cmdClone.Run()
	if err != nil {
		return err
	}

	// Clean Git files after cloning
	err = cleanGitFiles(destDir)
	if err != nil {
		return err
	}

	return nil
}

func cleanGitFiles(dir string) error {
	gitRelatedFiles := []string{
		".git",
		".gitattributes",
		".gitmodules",
	}

	for _, file := range gitRelatedFiles {
		err := os.RemoveAll(filepath.Join(dir, file))
		if err != nil {
			return err
		}
	}

	return nil
}
