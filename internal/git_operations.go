package internal

import (
	"fmt"
	"os"
	"strings"
	
	"github.com/go-git/go-git/v6/plumbing"
	git "github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing/object"
)

func GetCommitListWithResult() *AnalysisResult {
	result := NewAnalysisResult()
	
	gitCommitList, err := getGitCommitList()
	if err != nil {
		fmt.Printf("❌ Erreur lors de la récupération des commits: %v\n", err)
		os.Exit(1)
	}
	
	commitList := convertToCustomCommit(gitCommitList)
	
	for _, commit := range commitList {
		commit.DisplayQualityReport()
		result.AddCommit(commit)
	}
	
	result.CalculateGlobalScore()
	result.DisplayGlobalReport()
	
	return result
}

func GetCommitListOfBranchWithResult(branchName string) *AnalysisResult {
	result := NewAnalysisResult()
	
	gitCommitList, err := getGitCommitListFromBranch(branchName)
	if err != nil {
		fmt.Printf("❌ Erreur lors de la récupération des commits de la branche '%s': %v\n", branchName, err)
		os.Exit(1)
	}
	
	commitList := convertToCustomCommit(gitCommitList)
	
	for _, commit := range commitList {
		commit.DisplayQualityReport()
		result.AddCommit(commit)
	}
	
	result.CalculateGlobalScore()
	result.DisplayGlobalReport()
	
	return result
}

func getGitCommitList() ([]*object.Commit, error) {
	repo, err := git.PlainOpen(".")
	if err != nil {
		return nil, fmt.Errorf("impossible d'ouvrir le dépôt Git: %w", err)
	}
	
	ref, err := repo.Head()
	if err != nil {
		return nil, fmt.Errorf("impossible de récupérer HEAD: %w", err)
	}
	
	iter, err := repo.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		return nil, fmt.Errorf("impossible de récupérer l'historique: %w", err)
	}
	
	gitCommits := []*object.Commit{}
	err = iter.ForEach(func(c *object.Commit) error {
		gitCommits = append(gitCommits, c)
		return nil
	})
	
	if err != nil {
		return nil, fmt.Errorf("erreur lors du parcours des commits: %w", err)
	}
	
	return gitCommits, nil
}

func getGitCommitListFromBranch(branchName string) ([]*object.Commit, error) {
	repo, err := git.PlainOpen(".")
	if err != nil {
		return nil, fmt.Errorf("impossible d'ouvrir le dépôt Git: %w", err)
	}
	
	ref, err := repo.Reference(plumbing.ReferenceName("refs/heads/"+branchName), true)
	if err != nil {
		return nil, fmt.Errorf("branche '%s' introuvable: %w", branchName, err)
	}
	
	iter, err := repo.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		return nil, fmt.Errorf("impossible de récupérer l'historique de la branche '%s': %w", branchName, err)
	}
	
	gitCommits := []*object.Commit{}
	err = iter.ForEach(func(c *object.Commit) error {
		gitCommits = append(gitCommits, c)
		return nil
	})
	
	if err != nil {
		return nil, fmt.Errorf("erreur lors du parcours des commits de la branche '%s': %w", branchName, err)
	}
	
	return gitCommits, nil
}

func convertToCustomCommit(gitCommits []*object.Commit) []Commit {
	customCommitList := []Commit{}
	for _, gitCommit := range gitCommits {
		var commit Commit
		firstLine := strings.Split(gitCommit.Message, "\n")[0]
		commit.ParseHeader(firstLine)
		commit.ParseBodyAndFooter(gitCommit.Message)
		commit.CalculateQualityScore()
		customCommitList = append(customCommitList, commit)
	}
	return customCommitList
}