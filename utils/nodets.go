package utils

import (
	"github.com/spf13/cobra"
)

var genNodeTS = &cobra.Command{
	Use:   "ts-gen",
	Short: "Generate nodejs project using TypeScript",
	Long: `Generate nodejs project using Typescript
	Example usage:
	gno ts-gen -D ./projects/myapp -l "lib1 lib2" -d "lib3 lib4"`,
	Run: generateNodeTS,
}

func init() {
	rootCmd.AddCommand(genNodeTS)
	genNodeTS.Flags().StringP("directory", "D", "./myApp", "Output directory for the project")
	genNodeTS.Flags().BoolP("yes", "y", false, "Generate default NodeJs package.json file")
	genNodeTS.Flags().StringP("libs", "l", " ", "List of Node.js libraries to install")
	genNodeTS.Flags().StringP("dev-libs", "d", " ", "List of Node.js libraries to install")
}

func generateNodeTS(cmd *cobra.Command, args []string) {
	dir, _ := cmd.Flags().GetString("directory")
	createDirectory(dir)
	checkNPMInstallation()

	changeDirectory(dir)

	initNodeProject(cmd)

	generateTSConfigFile()

	installLibraries(cmd, "libs")
	installLibraries(cmd, "dev-libs")

	createDirectory("./src")
	createDirectory("./dist")

	InstallTSC()

	createFile("main.ts", "./src")

}
