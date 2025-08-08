package stats

import (
	"git_contribution_cli/utils"
	"git_contribution_cli/scan"
)


func Stats(email string) {
	commits := processRepos(email)
	printCommitsStats(commits)
}

func processRepos(email string) map[int]int {
	filePath := scan.GetDotFilePath()
	repos := utils.ParseFileLinesToSlices(filePath)
	daysInMap := 365

	commits := make(map[int]int, daysInMap)
	for i := daysInMap; i >0; i-- {
		commits[i] = 0
	}

	for _, path := range repos {
		commits = fillCommits(email, path, commits)
	}
	return commits
}

func fillCommits(email string, path string, commits map[int]int) map[int]int {
	return make(map[int]int)
}

func printCommitsStats(commits map[int]int) {
}
