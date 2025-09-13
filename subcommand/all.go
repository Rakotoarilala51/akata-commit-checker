package subcommand

import (
	"github.com/Rakotoarilala51/akata-commit-checker/internal"
	"github.com/spf13/cobra"
)

var AllCmd = &cobra.Command{
	Use: "all",
	Short: "Evaluation de tous les commits dans tous les branches",
	Run: func(cmd *cobra.Command, args []string) {
		result := internal.GetCommitListWithResult()
		result.ExitWithCode()
	},
}