package internal

import (
	"log"
	"strings"

	git "github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing/object"
)


func GetCommitList() {
	gitCommitList := getGitCommitList()
	commitList := convertToCustomCommit(gitCommitList)
	for _, commit := range commitList{
		commit.DisplayQualityReport()
	}
}

func getGitCommitList()[]*object.Commit{
	repo, err := git.PlainOpen(".")
	if err != nil {
		log.Fatalln(err)
	}
	ref, err := repo.Head()
	if err != nil{
		log.Fatalln(err)
	}
	iter, err := repo.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil{
		log.Fatalln(err)
	}
	gitCommits := []*object.Commit{}
	err = iter.ForEach(func(c *object.Commit) error {
		gitCommits = append(gitCommits, c)
		return nil
	})
	if err != nil{
		log.Fatalln(err)
	}
	return gitCommits
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

