package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all podcasts in your Spotify library",
	Long:  `Display a list of all podcasts currently in your Spotify library.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: Implement list functionality
		fmt.Println("List command called")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().String("format", "table", "Output format: table, json, csv")
	listCmd.Flags().Int("limit", 50, "Maximum number of podcasts to display")
}
