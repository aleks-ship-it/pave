package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/aleks-ship-it/pave/internal/linker"
)

var linkName string
var linkPath string

var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Create a symbolic link",
	Long:  `Create a symbolic link with the specified name and path.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if verbose {
			fmt.Printf("Linking %s -> %s (dry-run=%v)\n", linkName, linkPath, dryRun)
		}
		if dryRun {
			fmt.Printf("[DRY-RUN] Would link %s -> %s\n", linkName, linkPath)
			return nil
		}
		err := linker.CreateLink(linkName, linkPath, verbose)
		if err != nil {
			return fmt.Errorf("failed to create link: %w", err)
		}
		if verbose {
			fmt.Println("Link created successfully")
		}
		return nil
	},
}

func init() {
	linkCmd.Flags().StringVar(&linkName, "name", "", "Name of the link")
	linkCmd.Flags().StringVar(&linkPath, "path", "", "Path for the link")
	linkCmd.MarkFlagRequired("name")
	linkCmd.MarkFlagRequired("path")
}
