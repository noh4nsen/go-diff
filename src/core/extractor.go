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
		project := extractProject(dir)

		changes.ModifiedFiles = append(changes.ModifiedFiles, file)

		if !contains(changes.ModifiedDirs, dir) {
			changes.ModifiedDirs = append(changes.ModifiedDirs, dir)
		}
		if !contains(changes.ModifiedProjects, project) {
			changes.ModifiedProjects = append(changes.ModifiedProjects, project)
		}
	}

	return changes
}

func extractDirectory(file string) string {
	var path string
	index := strings.LastIndex(file, "/")
	if index > -1 {
		path = file[:index]
	} else {
		path = "."
	}
	return path
}

func extractProject(dir string) string {
	var path string
	index := strings.LastIndex(dir, "/")
	if index > -1 {
		path = dir[:index]
	} else {
		path = dir
	}
	return path
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}

	return false
}
