package stats

import (
	"git_contribution_cli/scan"
	"git_contribution_cli/utils"
	"time"

	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing/object"
)

const daysInLastYear = 365
const outOfRange = -1

func Stats(email string) {
	commits := processRepos(email)
	printCommitsStats(commits)
}

func processRepos(email string) map[int]int {
	filePath := scan.GetDotFilePath()
	repos := utils.ParseFileLinesToSlices(filePath)
	daysInMap := daysInLastYear

	commits := make(map[int]int, daysInMap)
	for i := daysInMap; i > 0; i-- {
		commits[i] = 0
	}

	for _, path := range repos {
		commits = fillCommits(email, path, commits)
	}
	return commits
}

// given a repo path, get the commits and put them in the commits map
func fillCommits(email string, path string, commits map[int]int) map[int]int {
	// making git repo from path
	repo, err := git.PlainOpen(path)
	if err != nil {
		panic(err)
	}
	// get HEAD ref
	ref, err := repo.Head()
	if err != nil {
		panic(err)
	}

	// get commit iterator
	iterator, err := repo.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		panic(err)
	}

	// build map
	offset := findOffset()
	err = iterator.ForEach(func(c *object.Commit) error {
		// getting days
		daysAgo := countDaysSinceDate(c.Author.When) + offset

		// compare email
		if c.Author.Email != email {
			return nil
		}

		// checking day range
		if daysAgo != outOfRange {
			commits[daysAgo]++
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return commits
}

func getBeginningOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	startOfDay := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	return startOfDay
}

func countDaysSinceDate(date time.Time) int {
	days := 0
	now := getBeginningOfDay(time.Now())
	for date.Before(now) {
		date = date.Add(time.Hour * 24)
		days++
		if days > daysInLastYear {
			return outOfRange
		}
	}
	return days
}

func findOffset() int {
    var offset int
    weekday := time.Now().Weekday()

    switch weekday {
    case time.Sunday:
        offset = 7
    case time.Monday:
        offset = 6
    case time.Tuesday:
        offset = 5
    case time.Wednesday:
        offset = 4
    case time.Thursday:
        offset = 3
    case time.Friday:
        offset = 2
    case time.Saturday:
        offset = 1
    }
    return offset
}

func printCommitsStats(commits map[int]int) {
}
