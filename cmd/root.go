package cmd

import (
	"orbitfs/cmd/add"
	"orbitfs/cmd/get"
	"orbitfs/cmd/list"
	"orbitfs/cmd/remove"

	"github.com/spf13/cobra"
)

func Execute() {
	var rootCmd = &cobra.Command{
		Use:   "orbitFS command",
		Short: "Central command to perform operations like: storing, fetching, listing, deleting etc.",
	}
	rootCmd.AddCommand(add.Command())
	rootCmd.AddCommand(get.Command())
	rootCmd.AddCommand(list.Command())
	rootCmd.AddCommand(remove.Command())

	rootCmd.Execute()
}
