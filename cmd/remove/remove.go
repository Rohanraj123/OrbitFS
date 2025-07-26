package remove

import (
	"github.com/spf13/cobra"
)

// RemoveOptions holds the flags for the remove command
type RemoveOptions struct {
	ID     string // ID or hash of the file to delete
	Force  bool   // Skip confirmation
	Forget bool   // Remove from local cache/index as well
	Unpin  bool   // Unpin from IPFS-like systems (if applicable)
	All    bool   // Remove all files (with confirmation)
	DryRun bool   // Show what would be deleted without deleting
}

func Command() *cobra.Command {
	opts := &RemoveOptions{}

	removeCmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove a file from OrbitFS or its index",
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO: Add actual removal logic using opts.ID, opts.Force, etc.
			return nil
		},
	}

	// Flags
	removeCmd.Flags().StringVarP(&opts.ID, "id", "i", "", "ID or hash of the file to remove (required)")
	removeCmd.Flags().BoolVarP(&opts.Force, "force", "f", false, "Force deletion without confirmation prompt")
	removeCmd.Flags().BoolVar(&opts.Forget, "forget", false, "Forget the file locally without affecting the network")
	removeCmd.Flags().BoolVar(&opts.Unpin, "unpin", false, "Unpin the file from the IPFS-like network")
	removeCmd.Flags().BoolVar(&opts.All, "all", false, "Remove all stored files (requires --force)")
	removeCmd.Flags().BoolVar(&opts.DryRun, "dry-run", false, "Simulate removal and print what would be deleted")

	// Mark ID as required unless --all is set
	removeCmd.MarkFlagsMutuallyExclusive("id", "all")
	removeCmd.MarkFlagsRequiredTogether("all", "force")

	return removeCmd
}
