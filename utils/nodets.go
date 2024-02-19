package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

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
	fmt.Println("Generating template...")

	// Clone repo
	err := cloneAndCheckoutWithLoading(repoUrl, destDir)
	if err != nil {
		fmt.Println("Error when generating template:", err)
		return
	}
	const totalSteps = 10
	for i := 0; i <= totalSteps; i++ {
		// Print progress bar
		const progressBarWidth = 50
		progress := i * progressBarWidth / totalSteps
		bar := strings.Repeat("=", progress) + strings.Repeat(" ", progressBarWidth-progress)
		fmt.Printf("\rProgress: [%s] %d%%", bar, (i*100)/totalSteps)

		os.Stdout.Sync()
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("\nTemplate generated successfully")
}

func cloneAndCheckoutWithLoading(repoURL, destDir string) error {
	// Create the destination directory if it doesn't exist
	CreateDirectoryTemp(destDir)

	// Clone the repository
	cmdClone := exec.Command("git", "clone", "-b", "main", "--single-branch", "--depth", "1", repoURL, destDir)

	// Redirect stdout and stderr
	var out bytes.Buffer
	cmdClone.Stdout = &out
	cmdClone.Stderr = &out

	// Start the clone command
	if err := cmdClone.Start(); err != nil {
		return err
	}

	// Wait for the command to finish
	if err := cmdClone.Wait(); err != nil {
		return err
	}

	// Clean Git files after cloning
	err := cleanGitFiles(destDir)
	if err != nil {
		return err
	}

	return nil
}
