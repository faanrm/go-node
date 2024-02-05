package test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/faanrm/go-node/utils"
	"github.com/spf13/cobra"
)

func TestCreateDirectory(t *testing.T) {
	dir := "testDir"
	defer os.RemoveAll(dir)
	utils.CreateDirectory(dir)

	// Check if the directory was created
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		t.Error("Expected directory to be created, but it does not exist")
	}
}

func TestCheckNPMInstallation(t *testing.T) {

	utils.CheckNPMInstallation()
}

func TestChangeDirectory(t *testing.T) {
	// Test with an existing directory
	currentDir, _ := os.Getwd()
	utils.ChangeDirectory(currentDir)
	newDir, _ := os.Getwd()

	if currentDir != newDir {
		t.Errorf("Expected directory to be changed to %s, got %s", currentDir, newDir)
	}

	// Test with a non-existing directory
	nonExistingDir := "nonExistingDir"
	err := utils.ChangeDirectory(nonExistingDir)
	if err == nil {
		t.Error("Expected error for changing to a non-existing directory, but got nil")
	}
}

func TestCreateFile(t *testing.T) {
	// Create a temporary directory for the test
	dir := t.TempDir()
	fmt.Printf("Test directory: %s\n", dir)

	fileName := "testFile.txt"
	content := "// Start writing here\nconsole.log(\"Hello World\");"

	// Create the file
	utils.CreateFile(fileName, dir)

	// Check if the file was created
	filePath := filepath.Join(dir, fileName)
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		t.Errorf("Expected file to be created, but it does not exist at path: %s", filePath)
		return
	}

	// Check the content of the file
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		t.Errorf("Error reading file: %v", err)
		return
	}

	if string(fileContent) != content {
		t.Errorf("Expected file content to be:\n%s\nGot:\n%s", content, string(fileContent))
	}
}

func TestInitNodeProject(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().Bool("yes", false, "")

	// Test with default flags
	utils.InitNodeProject(cmd)

	// Test with custom flags
	cmd.Flags().Set("yes", "true")
	utils.InitNodeProject(cmd)
}

func TestInstallLibraries(t *testing.T) {
	// Assuming cobra.Command and flags for testing
	cmd := &cobra.Command{}
	cmd.Flags().String("libs", "library1 library2", "")

	// Test with default flags
	utils.InstallLibraries(cmd, "libs")

	// Test with custom flags
	cmd.Flags().Set("libs", "library3 library4")
	utils.InstallLibraries(cmd, "libs")

}

func TestGenerateTSConfigFile(t *testing.T) {
	// Assuming the structure of TSConfig and expected JSON
	expectedJSON := `{"compilerOptions":{"target":"ES6","module":"CommonJS","outDir":"./dist","rootDir":"./src","strict":true,"esModuleInterop":true,"skipLibCheck":true},"include":["src/**/*.ts"],"exclude":["node_modules"]}`

	utils.GenerateTSConfigFile()
	// Check if the tsconfig.json file was generated
	_, err := os.Stat("tsconfig.json")
	if os.IsNotExist(err) {
		t.Error("Expected tsconfig.json file to be generated, but it does not exist")
	}

	// Check the content of the generated tsconfig.json file
	fileContent, _ := os.ReadFile("tsconfig.json")
	if string(fileContent) != expectedJSON {
		t.Errorf("Expected tsconfig.json content to be %s, got %s", expectedJSON, string(fileContent))
	}
}

func TestInstallTSC(t *testing.T) {
	// Assuming tsc is not installed for the testing environment
	utils.InstallTSC()

	// Check if tsc is installed after running InstallTSC
	cmd := exec.Command("tsc", "--version")
	err := cmd.Run()
	if err != nil {
		t.Error("Expected tsc to be installed, but it is not")
	}
}

func TestCreateDirectoryTemp(t *testing.T) {
	dir := "tempDir"
	defer os.RemoveAll(dir) // Cleanup

	utils.CreateDirectory(dir)

	// Check if the temporary directory was created
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		t.Error("Expected temporary directory to be created, but it does not exist")
	}
}

func TestMoveFiles(t *testing.T) {
	// Assuming source and destination directories for testing
	sourceDir := "sourceDir"
	destDir := "destDir"
	defer os.RemoveAll(destDir) // Cleanup

	// Create some test files in the source directory
	os.MkdirAll(sourceDir, 0755)
	os.Create(sourceDir + "/file1.txt")
	os.Create(sourceDir + "/file2.txt")

	utils.MoveFiles(nil, destDir, sourceDir)

	// Check if the files were moved to the destination directory
	_, err := os.Stat(destDir + "/file1.txt")
	if os.IsNotExist(err) {
		t.Error("Expected file1.txt to be moved to the destination directory, but it does not exist")
	}

	_, err = os.Stat(destDir + "/file2.txt")
	if os.IsNotExist(err) {
		t.Error("Expected file2.txt to be moved to the destination directory, but it does not exist")
	}

	// Check if the source directory was removed
	_, err = os.Stat(sourceDir)
	if !os.IsNotExist(err) {
		t.Error("Expected source directory to be removed, but it still exists")
	}
}

// Helper function to clean up after tests, if necessary
func cleanupTest() {
	// Remove any test-specific files or directories
	os.Remove("tsconfig.json")
}

// Run cleanup after all tests have been executed
func TestMain(m *testing.M) {
	code := m.Run()
	cleanupTest()
	os.Exit(code)
}
