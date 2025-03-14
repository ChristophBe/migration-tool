/*
Copyright © 2025 Christoph Becker <post@christopb.de>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [filename]",
	Short: "Add a file to the migration definition.",
	Long:  `Add a file to the migration definition. The file will be added to the end of the migration definition.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return acts.AddStepFile(baseFolder, args[0])
	},
}

func init() {
	addCmd.Args = cobra.ExactArgs(1)
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
