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
		dir := extractDirectory(file)

		changes.ModifiedFiles = append(changes.ModifiedFiles, file)

		if !contains(changes.ModifiedDirs, dir) {
			changes.ModifiedDirs = append(changes.ModifiedDirs, dir)
		}
	}

	return changes
}

func extractDirectory(file string) string {
	var dir string
	index := strings.LastIndex(file, "/")
	if index > -1 {
		dir = file[:index]
	} else {
		dir = "."
	}
	return dir
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}

	return false
}
