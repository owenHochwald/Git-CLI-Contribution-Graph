package main

import (
	"flag"
	"git_contribution_cli/scan"
	"git_contribution_cli/stats"
)


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

	stats.Stats(email)
	println()
}