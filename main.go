package main

import (
	"flag"
	"git_contribution_cli/scan"
)

// scan a path and crawl it including its subdirectories to find git repos


// generates the nice contributions graph
func stats(email string) {
	print("stats")
}

func main() {
	var folder string
	var email string

	flag.StringVar(&folder, "add", "", "add a folder to scan for git repos")
	flag.StringVar(&email, "email", "your@email.com", "the email to scan for")
	flag.Parse()

	if folder != "" {
		scan.Scan(folder)
		return
	}

	stats(email)
	println()
}