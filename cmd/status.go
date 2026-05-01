package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/aleks-ship-it/pave/internal/linker"
)

var statusName string

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show status of a symbolic link",
	Long:  `Show the status of a specific symbolic link or all links.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if statusName == "" {
			links, err := linker.ListLinks(verbose)
			if err != nil {
				return fmt.Errorf("failed to list links: %w", err)
			}
			if len(links) == 0 {
				fmt.Println("No links found")
				return nil
			}
			fmt.Println("Link status:")
			for _, link := range links {
				status := "valid"
				if link.Status != "" {
					status = link.Status
				}
				fmt.Printf("  %s: %s -> %s [%s]\n", link.Name, link.Path, link.Target, status)
			}
		} else {
			link, err := linker.StatusLink(statusName, verbose)
			if err != nil {
				return fmt.Errorf("failed to get status: %w", err)
			}
			if link == nil {
				fmt.Printf("Link %q not found\n", statusName)
				return nil
			}
			fmt.Printf("Name: %s\n", link.Name)
			fmt.Printf("Path: %s\n", link.Path)
			fmt.Printf("Target: %s\n", link.Target)
			fmt.Printf("Status: %s\n", link.Status)
		}
		return nil
	},
}

func init() {
	statusCmd.Flags().StringVar(&statusName, "name", "", "Name of the link to check")
}
