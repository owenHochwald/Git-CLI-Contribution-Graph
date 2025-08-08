package scan

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"
)

func Scan(folder string) {
	repos := recursiveScanFolder(folder)
	filePath := getDotFilePath()
	addNewSliceElementsToFile(filePath, repos)
	fmt.Printf("\n\nAdded %d new repos to scan\n", len(repos))
}

func recursiveScanFolder(folder string) []string {
	return scanGitFolders(make([]string, 0), folder)

}

func scanGitFolders(folders []string, folder string) []string {
	folder = strings.TrimSuffix(folder, "/") // safety to catch any trailing slashes

	// open folder
	f, err := os.Open(folder)
	if err != nil {
		log.Fatal(err)
	}
	// grab files
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	var path string

	for _, file := range files {
		if file.IsDir() {
			if file.Name() == ".git" {
				// trim .git stuff we added
				path = strings.TrimSuffix(path, ".git")
				fmt.Println(path)
				folders = append(folders, path)
				continue
			} else if file.Name() == "vendor" || file.Name() == "node_modules" {
				continue
			} else {
				// recursively scan subfolders
				path = folder + "/" + file.Name()
				folders = scanGitFolders(folders, path)
			}

		}
	}
	return folders

}

func getDotFilePath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	dotFile := usr.HomeDir + "/.gitstats"
	return dotFile

}

func addNewSliceElementsToFile(filePath string, repos []string) {

}
