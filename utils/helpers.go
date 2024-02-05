package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

func createDirectory(dir string) {
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

func checkNPMInstallation() {
	checkNPM := exec.Command("npm", "--version")
	npmVerOut, err := checkNPM.Output()
	if err != nil {
		fmt.Println("Can you please verify your node installation")
		log.Fatal(err)
	}
	fmt.Printf("npm version : %v \n", string(npmVerOut))
}

func changeDirectory(dir string) {
	if err := os.Chdir(dir); err != nil {
		log.Fatalf("Error changing directory: %v", err)
	}
}

func createFile(fileName string, dir string) {

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

func initNodeProject(cmd *cobra.Command) {
	def, _ := cmd.Flags().GetBool("yes")
	cmdInit := exec.Command("npm", "init")
	if def {
		cmdInit.Args = append(cmdInit.Args, "-y")
	}

	cmdInit.Stdin = os.Stdin
	cmdInit.Stdout = os.Stdout
	cmdInit.Stderr = os.Stderr

	if err := cmdInit.Run(); err != nil {
		log.Fatal("Error executing command:", err)
	}

	fmt.Println("Node JS project initialized...")
}

func installLibraries(cmd *cobra.Command, flag string) {
	usage := cmd.Use
	librariesStr, _ := cmd.Flags().GetString(flag)
	libraries := strings.Fields(librariesStr)
	if len(libraries) > 0 {
		fmt.Println("Installing node modules:", strings.Join(libraries, ", "))

		npmArgs := []string{"install"}
		if flag == "dev-libs" {
			npmArgs = append(npmArgs, "--save-dev")
		}

		if usage == "ts-node" {
			npmArgs = append(npmArgs, "typescript")
		}

		npmArgs = append(npmArgs, libraries...)

		installCmd := exec.Command("npm", npmArgs...)
		installCmd.Stdin = os.Stdin
		installCmd.Stdout = os.Stdout
		installCmd.Stderr = os.Stderr

		if err := installCmd.Run(); err != nil {
			log.Fatalf("Error installing node modules: %v", err)
		}
		fmt.Println("Node modules installed successfully.")
	}
}

func generateTSConfigFile() {
	// Create an instance of the TSConfig structure with your desired values.
	config := TSConfig{
		CompilerOptions: CompilerOptions{
			Target:          "ES6",
			Module:          "CommonJS",
			OutDir:          "./dist",
			RootDir:         "./src",
			Strict:          true,
			EsModuleInterop: true,
			SkipLibCheck:    true,
		},
		Include: []string{"src/**/*.ts"},
		Exclude: []string{"node_modules"},
	}
	jsonData, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Write the JSON data to a file.
	fileName := "tsconfig.json"
	err = os.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}

	fmt.Printf("Generated %s\n", fileName)
}

func InstallTSC() {
	// Check if tsc is installed by running a command.
	cmd := exec.Command("tsc", "--version")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		// tsc is not installed; let's install it.
		fmt.Println("tsc is not installed. Waiting for TypeScript installation...")
		installCmd := exec.Command("npm", "install", "-g", "typescript")
		installCmd.Stdout = os.Stdout
		installCmd.Stderr = os.Stderr

		if err := installCmd.Run(); err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		fmt.Println("TypeScript has been installed successfully.")
	} else {
		// tsc is already installed.
		fmt.Println("TypeScript is already installed.")
	}

}

func createDirectoryTemp(dir string) {
	// Create the directory if it doesn't exist
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
	}
}

func moveFiles(cmd *cobra.Command, destinationDir string, sourceDir string) {
	// Get all files and subdirectories in the source directory
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
