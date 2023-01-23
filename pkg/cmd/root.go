package cmd

import (
	"os"

	// pkgCCtx "github.com/nitroci/nitroci-core/pkg/core/contexts"
	"github.com/spf13/cobra"
)

// var ctxInput pkgCCtx.CoreContextBuilderInput

var rootCmd = &cobra.Command{
	Use:   "nitroci",
	Short: "NitroCI - Boost your software development",
	Long: `NitroCI - Boost your software development
	
The command line interface implementes devops practices 
and it is not tied to a particolar language or farmework.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },ls
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.SilenceErrors = false
	rootCmd.SilenceUsage = true
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func rootOnInitialize() {
	// path, _ := os.Getwd()
	// profile, _ := rootCmd.PersistentFlags().GetString("profile")
	// workspace, _ := rootCmd.PersistentFlags().GetInt("workspace")
	// verbose, _ := rootCmd.PersistentFlags().GetBool("verbose")
	// ctxInput = pkgCCtx.CoreContextBuilderInput{
	// 	WorkingDirectory: path,
	// 	Profile:          profile,
	// 	Environment:      "",
	// 	WorkspaceDepth:   workspace - 1,
	// 	Verbose:          verbose,
	// }
}

func init() {
	cobra.OnInitialize(rootOnInitialize)
	rootCmd.PersistentFlags().StringP("profile", "p", "default", "set a specific profile")
	rootCmd.PersistentFlags().IntP("workspace", "w", 1, "set current workspace")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "output verbose output")
}
