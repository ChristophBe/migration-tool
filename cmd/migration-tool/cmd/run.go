/*
Copyright © 2025 Christoph Becker <post@christopb.de>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs every script that were not previously executed.",
	Long: `Runs every script that were not previously executed.
Before running the scripts, the scripts are checked for unexecuted changes. 
In case of unexecuted changes, the script is not executed.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := acts.Run(baseFolder)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}

	},
}

func init() {
	rootCmd.AddCommand(runCmd)

}
