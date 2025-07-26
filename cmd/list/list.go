package list

import (
	"github.com/spf13/cobra"
)

// ListOptions holds all flag values for the list command.
type ListOptions struct {
	Tags     string // comma‑separated filter, e.g. "music,images"
	Private  bool   // show only private files
	Limit    int    // max number of results
	SortBy   string // name | size | date
	Reverse  bool   // reverse sort order
	Detailed bool   // show extra columns (hash, size, createdAt, etc.)
	NoHeader bool   // suppress table header
}

func Command() *cobra.Command {
	opts := &ListOptions{}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List files stored in OrbitFS",
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO: Replace with actual list logic that queries metadata store
			// and prints results using opts (e.g., opts.Tags, opts.Limit …).
			return nil
		},
	}

	// Flags
	listCmd.Flags().StringVarP(&opts.Tags, "tags", "t", "", "Filter by tags (comma‑separated)")
	listCmd.Flags().BoolVar(&opts.Private, "private", false, "Show only private files")
	listCmd.Flags().IntVarP(&opts.Limit, "limit", "l", 0, "Maximum number of files to display (0 = no limit)")
	listCmd.Flags().StringVar(&opts.SortBy, "sort", "date", "Sort by: name | size | date")
	listCmd.Flags().BoolVar(&opts.Reverse, "reverse", false, "Reverse the sort order")
	listCmd.Flags().BoolVarP(&opts.Detailed, "detailed", "d", false, "Show detailed output (hash, size, timestamps)")
	listCmd.Flags().BoolVar(&opts.NoHeader, "no-header", false, "Omit table header in output")

	return listCmd
}
