package internal

func giveScore(commit Commit) int{
	if commit.isValidCommit{
		return 0;
	}
	score := 5
	if commit.description == ""{
		score--;
	}
	if commit.footer == ""{
		score--;
	}
	if commit.porte == ""{
		score--;
	}
	return score
}