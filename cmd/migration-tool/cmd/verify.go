/*
Copyright © 2025 Christoph Becker <post@christopb.de>
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var ChangeDetectedError = errors.New("scripts have changed")

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify checks if the scripts have changed.",
	Long:  `With verify, the scripts are checked for unexecuted changes.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		anyChanged, err := acts.Verify(baseFolder)
		if err != nil {
			return err
		}
		if anyChanged {
			fmt.Println("The scripts have changed. Please run 'migration-tool regenerate-hashes' to recalculate the hashes incase the changes are intended.")
			return ChangeDetectedError
		} else {
			fmt.Println("All scripts are unchanged.")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(verifyCmd)
}
