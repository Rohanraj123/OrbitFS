package add

import (
	"github.com/spf13/cobra"
)

type FileAdderMetadata struct {
	FilePath    string
	FileName    string
	Tags        string
	Private     bool
	ExpiresIn   string
	Pin         bool
	Description string
	Encrypt     bool
	ChunkSize   int
	Compression string
}

func Command() *cobra.Command {
	fileMetadata := &FileAdderMetadata{}

	var addCmd = &cobra.Command{
		Use:   "add",
		Short: "Store files in decentralized nodes",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Your logic to handle file addition goes here
			return nil
		},
	}

	addCmd.Flags().StringVarP(&fileMetadata.FilePath, "file-path", "f", "", "Path to file that needs to be stored")
	addCmd.Flags().StringVar(&fileMetadata.FileName, "name", "", "Custom name (alias) for the file in OrbitFS. Defaults to file name")
	addCmd.Flags().StringVar(&fileMetadata.Tags, "tags", "", "Comma-separated list of tags for searching/grouping later")
	addCmd.Flags().BoolVar(&fileMetadata.Private, "private", false, "Marks file as private (e.g., encryption, hidden listing)")
	addCmd.Flags().StringVar(&fileMetadata.ExpiresIn, "expires-in", "", "Optional expiry duration like 1h, 24h, 7d")
	addCmd.Flags().BoolVar(&fileMetadata.Pin, "pin", false, "Pin the file to prevent garbage collection (like IPFS)")
	addCmd.Flags().StringVar(&fileMetadata.Description, "desc", "", "Description for the file (for metadata)")
	addCmd.Flags().BoolVar(&fileMetadata.Encrypt, "encrypt", false, "Encrypt the file before upload")
	addCmd.Flags().IntVar(&fileMetadata.ChunkSize, "chunk-size", 0, "Split file into chunks (in KB/MB), useful for large files")
	addCmd.Flags().StringVar(&fileMetadata.Compression, "compression", "", "Enable compression (none, gzip, zstd, etc.)")

	addCmd.MarkFlagRequired("file-path")

	return addCmd
}
