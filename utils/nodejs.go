package utils

import (
	"fmt"
	"os"
	"os/exec"

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
	// Modify the template command to accept an argument for the folder name
	template.Flags().StringP("directory", "D", "./myApp", "Output directory for the project")
	template.Flags().StringP("database", "d", "mongo", "Choose database type: mongo or sql")

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
	repoURLSQL := "https://github.com/faanrm/NodeJs-Sequelize-Starter"
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

	// Create the destination directory if it doesn't exist
	CreateDirectoryTemp(destDir)
	fmt.Println("Generating template ...")

	// Afficher un message de chargement avant de cloner le dépôt
	fmt.Println("waiting ...")
	// Indicateur de progression initialisé à 0%
	printProgress(0)

	// Clone du dépôt
	cmdClone := exec.Command("git", "clone", repoURL)
	cmdClone.Stdout = os.Stdout
	cmdClone.Stderr = os.Stderr

	// Rediriger la sortie standard et la sortie d'erreur vers /dev/null pour cacher les commandes git
	cmdClone.Stdout = nil
	cmdClone.Stderr = nil

	// Démarre le processus de clonage
	err := cmdClone.Start()
	if err != nil {
		fmt.Println("Erreur lors du démarrage du clonage:", err)
		return
	}

	err = cmdClone.Wait()
	if err != nil {
		fmt.Println("Erreur lors du clonage:", err)
		return
	}

	printProgress(100)

	if dbType == "mongo" {
		repoDir := "API-EXPRESS"
		MoveFiles(cmd, destDir, repoDir)
	} else if dbType == "sql" {
		repoDir := "NodeJs-Sequelize-Starter"
		MoveFiles(cmd, destDir, repoDir)
	}

	fmt.Println("\nTemplate generated successfully")
}

func printProgress(progress int) {
	fmt.Printf("\rChargement en cours... %d%%", progress)
}
