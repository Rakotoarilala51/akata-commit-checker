package main

import (
	"github.com/Rakotoarilala51/akata-commit-checker/internal"
	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command;

func main(){
	internal.GetCommitList()
	rootCmd.Execute();
}

func init(){
	rootCmd = &cobra.Command{
		Use: "check-commit",
		Short: "Outils d'évaluation de qualité de message de commit (selon la convention de commit d'AKATA GOAVANA) dans un dépot git \nHO GOAVANA HATRANY",
	}
}