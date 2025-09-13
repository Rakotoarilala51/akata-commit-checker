package internal

import (
	"fmt"
	"regexp"
	"strings"
)

type Commit struct{
	fullCommit string
	branch string
	types string
	porte string
	sujet string
	description string
	footer string
	isValidCommit bool
	score int
}
func (c *Commit) ParseBodyAndFooter(message string) {
	c.fullCommit = message
	commitMessage := strings.Split(message, "\n")
	if len(commitMessage) > 1 {
		footerStartIndex := -1
		for i := 1; i < len(commitMessage); i++ {
			line := strings.ToLower(strings.TrimSpace(commitMessage[i]))
			if strings.HasPrefix(line, "close") {
				footerStartIndex = i
				break
			}
		}
		if footerStartIndex != -1 {
			c.description = strings.Join(commitMessage[1:footerStartIndex], "\n")
			c.footer = strings.Join(commitMessage[footerStartIndex:], "\n")
		} else {
			c.description = strings.Join(commitMessage[1:], "\n")
		}
	}
}
func (commit *Commit) ParseHeader(message string){
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
func (c *Commit) CalculateQualityScore() {
	if !c.isValidCommit {
		c.score = 0
		return
	}
	score := 3
	if strings.TrimSpace(c.porte) != "" {
		score++
	}
	if strings.TrimSpace(c.description) != "" {
		score++
	}
	if strings.TrimSpace(c.footer) != "" && score < 5 {
		score++
	}
	if score > 5 {
		score = 5
	}
	c.score = score
}