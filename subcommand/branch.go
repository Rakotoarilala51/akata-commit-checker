package subcommand

import (
	"fmt"
	"os"
	
	"github.com/Rakotoarilala51/akata-commit-checker/internal"
	"github.com/spf13/cobra"
)

var BranchCmd = &cobra.Command{
	Use:   "branch <nom-branche>",
	Short: "Evaluation de tous les commits dans une branche spécifique",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		branchName := args[0]
		if branchName == "" {
			fmt.Println("❌ Erreur: Le nom de la branche ne peut pas être vide")
			os.Exit(1)
		}
		
		result := internal.GetCommitListOfBranchWithResult(branchName)
		result.ExitWithCode() 
	},
}