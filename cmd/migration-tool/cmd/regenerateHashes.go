/*
Copyright © 2025 Christoph Becker <post@christopb.de>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// regenerateHashesCmd represents the regenerateHashes command
var regenerateHashesCmd = &cobra.Command{
	Use:   "regenerate-hashes",
	Short: "This command recalculates the hashes of the scripts.",
	Long: `Regenerate hashes of the scripts. This is useful if the scripts have been intentional changed.
Be careful, this can lead to consistent behavior while executing the scripts. It might prevent the run command to execute scripts in some cases.
It is recommended that this is only used to recalculate the hashes for scripts that were not executed in any environment before.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Recalculating hashes...")
		err := acts.RecalculateHashes(baseFolder)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(regenerateHashesCmd)
}
