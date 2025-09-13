package main

import (
	"os"
	"github.com/Rakotoarilala51/akata-commit-checker/subcommand"
	"github.com/spf13/cobra"
)

var (
	qualityThreshold int
	verbose         bool
)

var rootCmd *cobra.Command

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd = &cobra.Command{
		Use:   "check-commit",
		Short: "Outils d'évaluation de qualité de message de commit (selon la convention de commit d'AKATA GOAVANA)",
		Long: `Un outil CLI pour évaluer la qualité des messages de commit Git selon les standards AKATA.
		
Exemples d'usage:
  check-commit all                    # Analyser tous les commits
  check-commit branch main            # Analyser la branche main
  check-commit all --threshold 4      # Seuil de qualité à 4/5
  check-commit all --verbose          # Mode verbeux`,
	}

	rootCmd.PersistentFlags().IntVarP(&qualityThreshold, "threshold", "t", 3, 
		"Seuil de qualité requis (1-5)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, 
		"Mode verbeux avec détails supplémentaires")
	subcommand.SetGlobalConfig(&qualityThreshold, &verbose)
	
	rootCmd.AddCommand(subcommand.AllCmd)
	rootCmd.AddCommand(subcommand.BranchCmd)
}