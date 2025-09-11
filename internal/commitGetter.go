package internal

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	git "github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing/object")


func GetCommitList() []Commit{
	repo, err := git.PlainOpen(".")
	commitList := []Commit{}
	if err != nil{
		log.Fatalln(err)
	}
	ref, err := repo.Head()
	if err != nil{
		log.Fatal(err)
	}

	iter, err := repo.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		log.Fatal(err)
	}
	
	err = iter.ForEach(func(c *object.Commit) error {
		var commit Commit;
		message := c.Message
		commitMessage := strings.Split(message, "\n")
		validateMainCommit(commitMessage[0], &commit)
		if len(commitMessage)>1{
			commit.description = strings.Join(commitMessage[1:], "\n") 
		}
		commitList = append(commitList, commit)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, c:= range commitList{
		fmt.Printf("%+v\n", c)
	}
	return commitList
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
