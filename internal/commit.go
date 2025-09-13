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
func (c *Commit) DisplayQualityReport() {
	fmt.Printf("=== RAPPORT DE QUALITÉ DU COMMIT ===\n")
	fmt.Printf("Sujet: %s\n", c.sujet)
	fmt.Printf("Score: %d/5\n", c.score)

	switch c.score {
	case 0:
		fmt.Printf("❌ COMMIT INVALIDE\n")
		fmt.Printf("   - Format incorrect (doit respecter: <type>(<portée>): <sujet>)\n")
		fmt.Printf("   - Types valides: build, ci, docs, feat, fix, perf, refactor, style, test\n")
		
	case 3:
		fmt.Printf("✅ COMMIT VALIDE - BASIQUE\n")
		fmt.Printf("   ✓ Type: %s\n", c.types)
		fmt.Printf("   ✓ Sujet: présent\n")
		
		missing := []string{}
		if strings.TrimSpace(c.porte) == "" {
			missing = append(missing, "portée (scope)")
		}
		if strings.TrimSpace(c.description) == "" {
			missing = append(missing, "description")
		}
		if strings.TrimSpace(c.footer) == "" {
			missing = append(missing, "footer")
		}
		
		if len(missing) > 0 {
			fmt.Printf("⚠️  AMÉLIORATIONS POSSIBLES:\n")
			for _, item := range missing {
				fmt.Printf("   - Ajouter %s\n", item)
			}
		}
		
	case 4:
		fmt.Printf("✅ BON COMMIT\n")
		fmt.Printf("   ✓ Type: %s\n", c.types)
		fmt.Printf("   ✓ Sujet: présent\n")
		
		if strings.TrimSpace(c.porte) != "" {
			fmt.Printf("   ✓ Portée: %s\n", c.porte)
		} else {
			fmt.Printf("   - Portée: manquante\n")
		}
		
		if strings.TrimSpace(c.description) != "" {
			fmt.Printf("   ✓ Description: présente\n")
		} else {
			fmt.Printf("   - Description: manquante\n")
		}
		
		if strings.TrimSpace(c.footer) != "" {
			fmt.Printf("   ✓ Footer: présent\n")
		} else {
			fmt.Printf("   - Footer: manquant\n")
		}
		
	case 5:
		fmt.Printf("🌟 EXCELLENT COMMIT\n")
		fmt.Printf("   ✓ Type: %s\n", c.types)
		fmt.Printf("   ✓ Sujet: présent\n")
		fmt.Printf("   ✓ Portée: %s\n", c.porte)
		fmt.Printf("   ✓ Description: présente\n")
		
		if strings.TrimSpace(c.footer) != "" {
			fmt.Printf("   ✓ Footer: présent\n")
		}
		
		fmt.Printf("   🎉 Respect parfait des conventions Git!\n")
	}
	
	fmt.Printf("=====================================\n\n")
}
func (ar *AnalysisResult) SetThreshold(threshold int) {
    if threshold >= 1 && threshold <= 5 {
        ar.QualityThreshold = threshold
    }
}