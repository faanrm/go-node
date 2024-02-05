package utils

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var templateTs = &cobra.Command{
	Use:   "template-ts",
	Short: "Generate node ts project using Typescript²	",
	Run:   generateNodeTsTemplate,
}
var genNodeTS = &cobra.Command{
	Use:   "ts-gen",
	Short: "Generate nodejs project using TypeScript",
	Long: `Generate nodejs project using Typescript
	Example usage:
	gno ts-gen -D ./projects/myapp -l "lib1 lib2" -d "lib3 lib4"`,
	Run: generateNodeTS,
}

func init() {
	rootCmd.AddCommand(templateTs)
	rootCmd.AddCommand(genNodeTS)
	genNodeTS.Flags().StringP("directory", "D", "./myApp", "Output directory for the project")
	genNodeTS.Flags().BoolP("yes", "y", false, "Generate default NodeJs package.json file")
	genNodeTS.Flags().StringP("libs", "l", " ", "List of Node.js libraries to install")
	genNodeTS.Flags().StringP("dev-libs", "d", " ", "List of Node.js libraries to install")
}
func generateNodeTsTemplate(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("Error: Please provide a folder name for the template.")
		return
	}
	repoUrl := "https://github.com/faanrm/nodeTs-template"
	destDir := args[0] //this to get the folder name from command line
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
func generateNodeTS(cmd *cobra.Command, args []string) {
	dir, _ := cmd.Flags().GetString("directory")
	CreateDirectory(dir)
	CheckNPMInstallation()

	ChangeDirectory(dir)

	InitNodeProject(cmd)

	GenerateTSConfigFile()

	InstallLibraries(cmd, "libs")
	InstallLibraries(cmd, "dev-libs")

	CreateDirectory("./src")
	CreateDirectory("./dist")

	InstallTSC()

	CreateFile("main.ts", "./src")

}
