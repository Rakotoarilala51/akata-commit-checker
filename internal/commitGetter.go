package internal

import (
	"fmt"
	"os"
)
var commitList []Commit

func GetCommitList(){
	commitPath:= ".git/logs/refs/heads/main"
	data, err := os.ReadFile(commitPath)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
