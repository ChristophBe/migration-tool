/*
Copyright © 2025 Christoph Becker <post@christopb.de>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify checks if the scripts have changed.",
	Long:  `With verify, the scripts are checked for unexecuted changes.`,
	Run: func(cmd *cobra.Command, args []string) {
		act := initActions(executionLogFile)
		anyChanged, err := act.Verify(baseFolder)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}
		if anyChanged {
			fmt.Println("The scripts have changed. Please run 'migration-tool regenerate-hashes' to recalculate the hashes incase the changes are intended.")
			os.Exit(1)
		} else {
			fmt.Println("All scripts are unchanged.")
		}
	},
}

func init() {
	rootCmd.AddCommand(verifyCmd)
}
