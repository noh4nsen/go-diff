package model

type Changes struct {
	ModifiedFiles    []string `json:"modifiedFiles"`
	ModifiedDirs     []string `json:"modifiedDirs"`
	ModifiedProjects []string `json:"modifiedProjects"`
}
