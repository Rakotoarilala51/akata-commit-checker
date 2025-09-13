package internal

import (
	"fmt"
	"os"
)

type AnalysisResult struct {
	TotalCommits    int
	ValidCommits    int
	InvalidCommits  int
	AverageScore    float64
	GlobalScore     int
	Commits         []Commit
	QualityThreshold int
}

func NewAnalysisResult() *AnalysisResult {
	return &AnalysisResult{
		QualityThreshold: 3,
	}
}

func (ar *AnalysisResult) AddCommit(commit Commit) {
	ar.Commits = append(ar.Commits, commit)
	ar.TotalCommits++
	
	if commit.isValidCommit {
		ar.ValidCommits++
	} else {
		ar.InvalidCommits++
	}
}

func (ar *AnalysisResult) CalculateGlobalScore() {
	if ar.TotalCommits == 0 {
		ar.GlobalScore = 0
		ar.AverageScore = 0
		return
	}

	totalScore := 0
	for _, commit := range ar.Commits {
		totalScore += commit.score
	}

	ar.AverageScore = float64(totalScore) / float64(ar.TotalCommits)
	ar.GlobalScore = int(ar.AverageScore)
}

func (ar *AnalysisResult) DisplayGlobalReport() {
	fmt.Printf("\n🔍 === RAPPORT GLOBAL D'ANALYSE ===\n")
	fmt.Printf("📊 Commits analysés: %d\n", ar.TotalCommits)
	fmt.Printf("✅ Commits valides: %d\n", ar.ValidCommits)
	fmt.Printf("❌ Commits invalides: %d\n", ar.InvalidCommits)
	fmt.Printf("📈 Score moyen: %.2f/5\n", ar.AverageScore)
	fmt.Printf("🎯 Score global: %d/5\n", ar.GlobalScore)
	
	// Status global
	if ar.GlobalScore >= ar.QualityThreshold {
		fmt.Printf("✅ QUALITÉ ACCEPTABLE (seuil: %d/5)\n", ar.QualityThreshold)
	} else {
		fmt.Printf("❌ QUALITÉ INSUFFISANTE (seuil: %d/5)\n", ar.QualityThreshold)
	}
	
	fmt.Printf("=====================================\n\n")
}

func (ar *AnalysisResult) GetExitCode() int {
	if ar.GlobalScore >= ar.QualityThreshold && ar.InvalidCommits == 0 {
		return 0 
	}
	return 1
}

func (ar *AnalysisResult) ExitWithCode() {
	exitCode := ar.GetExitCode()
	if exitCode == 0 {
		fmt.Printf("🎉 ANALYSE RÉUSSIE - Code de sortie: %d\n", exitCode)
	} else {
		fmt.Printf("⚠️  ANALYSE ÉCHOUÉE - Code de sortie: %d\n", exitCode)
	}
	os.Exit(exitCode)
}