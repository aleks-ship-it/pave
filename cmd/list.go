package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/aleks-ship-it/pave/internal/linker"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all symbolic links",
	Long:  `List all managed symbolic links.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		links, err := linker.ListLinks(verbose)
		if err != nil {
			return fmt.Errorf("failed to list links: %w", err)
		}
		if len(links) == 0 {
			fmt.Println("No links found")
			return nil
		}
		fmt.Println("Managed links:")
		for _, link := range links {
			status := "valid"
			if link.Status != "" {
				status = link.Status
			}
			fmt.Printf("  %s -> %s [%s]\n", link.Name, link.Path, status)
		}
		return nil
	},
}
