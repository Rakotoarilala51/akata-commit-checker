package internal

import (
	"fmt"
	"os"
	"github.com/fatih/color"
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
	fmt.Printf("\n")
	
	green := color.New(color.FgGreen)
	green.Println("> Initializing commit analysis...")
	green.Println("> Scanning repository structure...")
	green.Println("> Processing commit metadata...")
	
	fmt.Printf("\n")
	fmt.Printf("╔══════════════════════════════════════════════════╗\n")
	fmt.Printf("║                COMMIT ANALYSIS RESULTS           ║\n")
	fmt.Printf("╠══════════════════════════════════════════════════╣\n")
	fmt.Printf("║                                                  ║\n")
	
	// Les couleurs n'affectent plus l'alignement !
	yellow := color.New(color.FgYellow)
	greenColor := color.New(color.FgGreen)
	red := color.New(color.FgRed)
	cyan := color.New(color.FgCyan)
	bold := color.New(color.Bold)
	
	fmt.Printf("║  Total commits............: ")
	yellow.Printf("%04d", ar.TotalCommits)
	fmt.Printf("                 ║\n")
	
	fmt.Printf("║  Valid commits............: ")
	greenColor.Printf("%04d", ar.ValidCommits)
	fmt.Printf("                 ║\n")
	
	fmt.Printf("║  Invalid commits..........: ")
	red.Printf("%04d", ar.InvalidCommits)
	fmt.Printf("                 ║\n")
	
	fmt.Printf("║  Average quality..........: ")
	cyan.Printf("%.2f/5", ar.AverageScore)
	fmt.Printf("               ║\n")
	
	fmt.Printf("║  Repository score.........: ")
	bold.Printf("%d/5", ar.GlobalScore)
	fmt.Printf("                  ║\n")
	
	fmt.Printf("║                                                  ║\n")
	
	if ar.GlobalScore >= ar.QualityThreshold {
		fmt.Printf("║  Status: [")
		greenColor.Printf(" PASS ")
		fmt.Printf("] Quality threshold met          ║\n")
		
		fmt.Printf("║  Required threshold.......: ")
		greenColor.Printf("%d/5", ar.QualityThreshold)
		fmt.Printf("                  ║\n")
	} else {
		fmt.Printf("║  Status: [")
		red.Printf(" FAIL ")
		fmt.Printf("] Quality below threshold   ║\n")
		
		fmt.Printf("║  Required threshold.......: ")
		red.Printf("%d/5", ar.QualityThreshold)
		fmt.Printf("               ║\n")
	}
	
	fmt.Printf("║                                                  ║\n")
	fmt.Printf("╚══════════════════════════════════════════════════╝\n")
	fmt.Printf("\n")
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