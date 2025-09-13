package main

import (
	"os"

	"github.com/Rakotoarilala51/akata-commit-checker/subcommand"
	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command;

func main(){
	if err:=rootCmd.Execute(); err != nil{
		os.Exit(1)
	}
}

func init(){
	rootCmd = &cobra.Command{
		Use: "check-commit",
		Short: "Outils d'évaluation de qualité de message de commit (selon la convention de commit d'AKATA GOAVANA) dans un dépot git \nHO GOAVANA HATRANY",
	}
	rootCmd.AddCommand(subcommand.AllCmd)
	rootCmd.AddCommand(subcommand.BranchCmd)
}