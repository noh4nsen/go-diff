package core

import (
	"go-diff/model"
	"strings"
)

func ExtractChanges(output string) model.Changes {
	changes := model.Changes{}
	files := strings.Split(output, "\n")

	for _, file := range files {
		if file == "" {
			continue
		}
		dir := file[:strings.LastIndex(file, "/")]
		changes.ModifiedFiles = append(changes.ModifiedFiles, file)

		if !contains(changes.ModifiedDirs, dir) {
			changes.ModifiedDirs = append(changes.ModifiedDirs, dir)
		}
	}

	return changes
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}

	return false
}
