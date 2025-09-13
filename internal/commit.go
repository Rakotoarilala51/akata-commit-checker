package internal

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/fatih/color"
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
func (c *Commit) DisplayQualityReport() {
	// Définition des couleurs
	red := color.New(color.FgRed)
	green := color.New(color.FgGreen)
	yellow := color.New(color.FgYellow)
	//bold := color.New(color.Bold)
	boldGreen := color.New(color.FgGreen, color.Bold)
	
	fmt.Printf("╔════════════════════════════════════════════════════════════╗\n")
	fmt.Printf("║                    COMMIT QUALITY ANALYSIS                 ║\n")
	fmt.Printf("╠════════════════════════════════════════════════════════════╣\n")
	fmt.Printf("║                                                            ║\n")
	fmt.Printf("║  Subject: %-47s ║\n", c.sujet)
	fmt.Printf("║  Quality Score: %d/5                                      ║\n", c.score)
	fmt.Printf("║                                                            ║\n")
	fmt.Printf("╠════════════════════════════════════════════════════════════╣\n")

	switch c.score {
	case 0:
		fmt.Printf("║  Status: ")
		red.Printf("INVALID COMMIT FORMAT")
		fmt.Printf("                          ║\n")
		fmt.Printf("║                                                            ║\n")
		fmt.Printf("║  CRITICAL ERRORS DETECTED:                                ║\n")
		fmt.Printf("║  > Format violation: <type>(<scope>): <subject>           ║\n")
		fmt.Printf("║  > Valid types: build, ci, docs, feat, fix, perf,         ║\n")
		fmt.Printf("║                 refactor, style, test                     ║\n")

	case 3:
		fmt.Printf("║  Status: ")
		yellow.Printf("BASIC VALID COMMIT")
		fmt.Printf("                             ║\n")
		fmt.Printf("║                                                            ║\n")
		fmt.Printf("║  PRESENT ELEMENTS:                                        ║\n")
		fmt.Printf("║  [+] Type: %-47s ║\n", c.types)
		fmt.Printf("║  [+] Subject: present                                      ║\n")
		
		missing := []string{}
		if strings.TrimSpace(c.porte) == "" {
			missing = append(missing, "scope")
		}
		if strings.TrimSpace(c.description) == "" {
			missing = append(missing, "description")
		}
		if strings.TrimSpace(c.footer) == "" {
			missing = append(missing, "footer")
		}
		
		if len(missing) > 0 {
			fmt.Printf("║                                                            ║\n")
			fmt.Printf("║  OPTIMIZATION OPPORTUNITIES:                              ║\n")
			for _, item := range missing {
				fmt.Printf("║  [-] Missing: %-43s ║\n", item)
			}
		}

	case 4:
		fmt.Printf("║  Status: ")
		green.Printf("GOOD COMMIT QUALITY")
		fmt.Printf("                            ║\n")
		fmt.Printf("║                                                            ║\n")
		fmt.Printf("║  ELEMENT ANALYSIS:                                        ║\n")
		fmt.Printf("║  [+] Type: %-47s ║\n", c.types)
		fmt.Printf("║  [+] Subject: present                                      ║\n")
		
		if strings.TrimSpace(c.porte) != "" {
			fmt.Printf("║  [+] Scope: %-46s ║\n", c.porte)
		} else {
			fmt.Printf("║  [-] Scope: missing                                        ║\n")
		}
		
		if strings.TrimSpace(c.description) != "" {
			fmt.Printf("║  [+] Description: present                                  ║\n")
		} else {
			fmt.Printf("║  [-] Description: missing                                  ║\n")
		}
		
		if strings.TrimSpace(c.footer) != "" {
			fmt.Printf("║  [+] Footer: present                                       ║\n")
		} else {
			fmt.Printf("║  [-] Footer: missing                                       ║\n")
		}

	case 5:
		fmt.Printf("║  Status: ")
		boldGreen.Printf("EXCELLENT COMMIT QUALITY")
		fmt.Printf("                       ║\n")
		fmt.Printf("║                                                            ║\n")
		fmt.Printf("║  PERFECT COMPLIANCE ACHIEVED:                             ║\n")
		fmt.Printf("║  [+] Type: %-47s ║\n", c.types)
		fmt.Printf("║  [+] Subject: present                                      ║\n")
		fmt.Printf("║  [+] Scope: %-46s ║\n", c.porte)
		fmt.Printf("║  [+] Description: present                                  ║\n")
		if strings.TrimSpace(c.footer) != "" {
			fmt.Printf("║  [+] Footer: present                                       ║\n")
		}
		fmt.Printf("║                                                            ║\n")
		fmt.Printf("║  ")
		green.Printf(">>> GIT CONVENTIONS: FULLY COMPLIANT <<<")
		fmt.Printf("              ║\n")
	}
	
	fmt.Printf("║                                                            ║\n")
	fmt.Printf("╚════════════════════════════════════════════════════════════╝\n")
	fmt.Printf("\n")
}


func (ar *AnalysisResult) SetThreshold(threshold int) {
    if threshold >= 1 && threshold <= 5 {
        ar.QualityThreshold = threshold
    }
}