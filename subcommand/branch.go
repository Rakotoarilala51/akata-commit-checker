package subcommand

import (
	"log"

	"github.com/Rakotoarilala51/akata-commit-checker/internal"
	"github.com/spf13/cobra"
)

var BranchCmd = &cobra.Command{
	Use: "branch",
	Short: "Evaluation de tous les commits dans un branche spécifiquse",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args)<=0{
			log.Fatalln("vous devrais ajouter le nom de Branche")
		}else{
			internal.GetCommitListOfBranch(args[0])
		}
		
	},
}