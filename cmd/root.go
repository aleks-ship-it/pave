package cmd

import (
	"github.com/spf13/cobra"
)

var (
	version = "dev"
	verbose bool
	dryRun  bool
)

var rootCmd = &cobra.Command{
	Use:   "pave",
	Short: "Pave - A CLI tool for managing links and generating installers",
	Long:  `Pave is a CLI tool that helps manage symbolic links and generate installation scripts.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "Enable verbose output")
	rootCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "Preview changes without applying them")
	rootCmd.Flags().BoolP("version", "V", false, "Print version information")

	rootCmd.SetVersionTemplate("{{.Version}}\n")
	rootCmd.Version = version

	rootCmd.AddCommand(linkCmd)
	rootCmd.AddCommand(unlinkCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(generateCmd)
}
