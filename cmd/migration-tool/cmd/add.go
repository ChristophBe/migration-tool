/*
Copyright © 2025 Christoph Becker <post@christopb.de>
*/
package cmd

import (
	"github.com/ChristophBe/migration-tool/internal/utils"
	"github.com/spf13/cobra"
)

const descriptionFlag = "description"
const descriptionFlagShort = "d"

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [filename]",
	Short: "Add a file to the migration definition.",
	Long: `Add a file to the migration definition. The file will be added to the end of the migration definition.
Only files in the same folder as the migration.yaml file or in a subfolder of this folder can be added.`,
	Args: func(cmd *cobra.Command, args []string) error {

		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return err
		}

		filename := args[0]

		if err := utils.FileInBaseFolderCheck(baseFolder, filename); err != nil {
			return err
		}

		if err := utils.FileExistsCheck(filename); err != nil {
			return err
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {

		description, err := cmd.Flags().GetString(descriptionFlag)
		if err != nil {
			return err
		}
		return acts.AddStepFile(baseFolder, args[0], description)
	},
}

func init() {

	addCmd.Flags().StringP(descriptionFlag, descriptionFlagShort, "", "Description of the step")
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
