package internal

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	git "github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing/object")


func GetCommitList() []Commit{
	gitCommitList := getGitCommitList()
	commitList := convertToCustomCommit(gitCommitList)
	for _, c:= range commitList{
		fmt.Printf("%+v\n", c)
	}
	return commitList
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
func convertToCustomCommit(gitCommits []*object.Commit) []Commit{
	customCommitList := []Commit{}
	for _, gitCommit := range gitCommits{
		var commit Commit
		commit.fullCommit = gitCommit.Message
		commitMessage := strings.Split(commit.fullCommit, "\n");
		validateMainCommit(commitMessage[0], &commit)
		if len(commitMessage)>1{
			commit.description = strings.Join(commitMessage[1:], "\n") 
		}
		customCommitList = append(customCommitList, commit)
	}
	return customCommitList

}
func validateMainCommit(message string, commit *Commit){
	validTypes := []string{
		"build", "ci", "docs", "feat", "fix", "perf", "refactor", "style", "test",
	}
	
	typePattern := fmt.Sprintf("(%s)", strings.Join(validTypes, "|"))
	regexPattern := fmt.Sprintf("^%s(?:\\(([^)]+)\\))?\\s*:\\s*(.+)$", typePattern)
	regex := regexp.MustCompile(regexPattern)
	
	matches := regex.FindStringSubmatch(message)
	commit.isValidCommit = matches!=nil
	if commit.isValidCommit{
		commit.types=matches[1]
		commit.porte=matches[2]
		commit.sujet=matches[3]
	}
}
