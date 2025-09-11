package main

import "github.com/spf13/cobra"

var rootCmd *cobra.Command;

func main(){
	rootCmd.Execute();
}

func init(){
	rootCmd = &cobra.Command{
		Use: "check-commit",
		Short: "Outils d'évaluation de qualité de message de commit (selon la convention de commit d'AKATA GOAVANA) dans un dépot git \nHO GOAVANA HATRANY",
	}
}