package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/aleks-ship-it/pave/internal/linker"
)

var unlinkName string

var unlinkCmd = &cobra.Command{
	Use:   "unlink",
	Short: "Remove a symbolic link",
	Long:  `Remove a symbolic link by name.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if verbose {
			fmt.Printf("Unlinking %s (dry-run=%v)\n", unlinkName, dryRun)
		}
		if dryRun {
			fmt.Printf("[DRY-RUN] Would unlink %s\n", unlinkName)
			return nil
		}
		err := linker.RemoveLink(unlinkName, verbose)
		if err != nil {
			return fmt.Errorf("failed to remove link: %w", err)
		}
		if verbose {
			fmt.Println("Link removed successfully")
		}
		return nil
	},
}

func init() {
	unlinkCmd.Flags().StringVar(&unlinkName, "name", "", "Name of the link to remove")
	unlinkCmd.MarkFlagRequired("name")
}
