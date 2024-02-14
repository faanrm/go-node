package utils

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func CreateDirectory(dir string) {
	if dir == "" {
		return
	}

	var stderr bytes.Buffer
	cmd := exec.Command("mkdir", dir)
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Error executing command:", err)
		fmt.Println("Standard error output:", stderr.String())
		return
	}

	fmt.Printf("Directory %v created successfully\n", dir)
}

/*
	func ChangeDirectory(dir string) {
		if err := os.Chdir(dir); err != nil {
			log.Fatalf("Error changing directory: %v", err)
		}
	}
*/
func ChangeDirectory(dir string) error {
	if err := os.Chdir(dir); err != nil {
		return fmt.Errorf("error changing directory: %v", err)
	}
	return nil
}
func CreateFile(fileName string, dir string) {

	if err := os.Chdir(dir); err != nil {
		log.Fatalf("Error changing directory: %v", err)
	}

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	content := `//Start writing here\n
	console.log("Hello World");`
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Printf("Successfully created and wrote to %s/%s\n", dir, fileName)

}

func CreateDirectoryTemp(dir string) {
	// Create the directory if it doesn't exist
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
	}
}

func MoveFiles(cmd *cobra.Command, destinationDir string, sourceDir string) {
	files, err := os.ReadDir(sourceDir)
	if err != nil {
		fmt.Println("Error reading source directory:", err)
		return
	}

	// Move each file to the destination directory
	for _, file := range files {
		sourcePath := fmt.Sprintf("%s/%s", sourceDir, file.Name())
		destinationPath := fmt.Sprintf("%s/%s", destinationDir, file.Name())

		// Move the file
		err := os.Rename(sourcePath, destinationPath)
		if err != nil {
			fmt.Printf("Error moving file %s to %s: %v\n", sourcePath, destinationPath, err)
		}
	}

	// Remove the empty source directory
	err = os.Remove(sourceDir)
	if err != nil {
		fmt.Println("Error removing source directory:", err)
	}
}
