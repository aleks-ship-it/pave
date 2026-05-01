package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/aleks-ship-it/pave/internal/generator"
)

var genName string
var genRepo string
var genBin string
var genOut string

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate installation scripts",
	Long:  `Generate installation scripts for the specified repository.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if verbose {
			fmt.Printf("Generating install scripts for %s (repo=%s, bin=%s, out=%s, dry-run=%v)\n",
				genName, genRepo, genBin, genOut, dryRun)
		}
		cfg := generator.Config{
			Name: genName,
			Repo: genRepo,
			Bin:  genBin,
		}
		if genOut == "" {
			genOut = "."
		}
		err := generator.Generate(cfg, genOut, dryRun, verbose)
		if err != nil {
			return fmt.Errorf("failed to generate scripts: %w", err)
		}
		return nil
	},
}

func init() {
	generateCmd.Flags().StringVar(&genName, "name", "", "Name of the tool")
	generateCmd.Flags().StringVar(&genRepo, "repo", "", "GitHub repository (owner/repo)")
	generateCmd.Flags().StringVar(&genBin, "bin", "", "Binary name")
	generateCmd.Flags().StringVar(&genOut, "out", "", "Output directory")
	generateCmd.MarkFlagRequired("name")
	generateCmd.MarkFlagRequired("repo")
	generateCmd.MarkFlagRequired("bin")
}
