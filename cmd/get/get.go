package get

import (
	"github.com/spf13/cobra"
)

type FileFetcherMetadata struct {
	ID           string
	Name         string
	OutputPath   string
	Decrypt      bool
	Verify       bool
	Version      string
	Uncompress   bool
	OutputFormat string
}

func Command() *cobra.Command {
	fileMetadata := &FileFetcherMetadata{}

	var getCmd = &cobra.Command{
		Use:   "get",
		Short: "Retrieve a file from decentralized storage",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Your logic to fetch the file goes here
			return nil
		},
	}

	getCmd.Flags().StringVarP(&fileMetadata.ID, "id", "i", "", "Unique identifier (CID/hash) of the file to fetch")
	getCmd.Flags().StringVarP(&fileMetadata.Name, "name", "n", "", "Alias/name of the file if CID is not used")
	getCmd.Flags().StringVarP(&fileMetadata.OutputPath, "output", "o", "", "Path to store the fetched file")
	getCmd.Flags().BoolVar(&fileMetadata.Decrypt, "decrypt", false, "Decrypt the file if it was stored encrypted")
	getCmd.Flags().BoolVar(&fileMetadata.Verify, "verify", false, "Verify the integrity of the downloaded file")
	getCmd.Flags().StringVar(&fileMetadata.Version, "version", "", "Fetch a specific version of the file")
	getCmd.Flags().BoolVar(&fileMetadata.Uncompress, "uncompress", false, "Uncompress the file after fetching")
	getCmd.Flags().StringVar(&fileMetadata.OutputFormat, "output-format", "", "Convert the file to a specific format (e.g., txt, json)")

	// Require at least one identifier: either ID or Name
	getCmd.MarkFlagsMutuallyExclusive("id", "name")

	return getCmd
}
