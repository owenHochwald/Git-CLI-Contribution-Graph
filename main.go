package main

import (
	"flag"
	"fmt"
	"git_contribution_cli/scan"
	"git_contribution_cli/stats"
	"git_contribution_cli/utils"
)

const usage = `GitStats - A tool to visualize your git contributions across multiple repositories

Usage:
  gitstats [flags]

Flags:
  -add string    Add a folder to scan recursively for git repositories
  -email string  The email to scan for in commit history (default: your@email.com)
  -list          List all currently tracked repositories
  -remove string Remove a repository from tracking
  -version       Show version information`

const version = "1.0.0"

func main() {
	var folder string
	var email string
	var showVersion bool
	var listRepos bool
	var removeRepo string

	flag.StringVar(&folder, "add", "", "add a folder to scan for git repos")
	flag.StringVar(&email, "email", "your@email.com", "the email to scan for")
	flag.BoolVar(&showVersion, "version", false, "show version information")
	flag.BoolVar(&listRepos, "list", false, "list all tracked repositories")
	flag.StringVar(&removeRepo, "remove", "", "remove a repository from tracking")

	flag.Usage = func() {
		fmt.Println(usage)
	}

	flag.Parse()

	if showVersion {
		fmt.Printf("gitstats version %s\n", version)
		return
	}

	if listRepos {
		repos := utils.ParseFileLinesToSlices(scan.GetDotFilePath())
		fmt.Println("Currently tracked repositories:")
		for _, repo := range repos {
			fmt.Printf("  %s\n", repo)
		}
		return
	}

	if removeRepo != "" {
		// TODO: Implement repository removal
		fmt.Printf("TODO: Remove repository: %s\n", removeRepo)
		return
	}

	if folder != "" {
		scan.Scan(folder)
		return
	}

	stats.Stats(email)
	println()
}
