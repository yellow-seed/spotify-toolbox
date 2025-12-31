package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete podcasts from your Spotify library",
	Long: `Delete podcasts from your Spotify library in bulk.
You can specify filters to select which podcasts to delete.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: Implement delete functionality
		fmt.Println("Delete command called")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().Bool("all", false, "Delete all podcasts")
	deleteCmd.Flags().String("filter", "", "Filter podcasts by name pattern")
	deleteCmd.Flags().Bool("dry-run", false, "Show what would be deleted without actually deleting")
}
