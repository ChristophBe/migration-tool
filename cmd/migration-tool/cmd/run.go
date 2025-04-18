/*
Copyright © 2025 Christoph Becker <post@christopb.de>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs every script that were not previously executed.",
	Long: `Runs every script that were not previously executed.
Before running the scripts, the scripts are checked for unexecuted changes. 
In case of unexecuted changes, the script is not executed.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return acts.Run(baseFolder)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

}
