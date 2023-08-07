package validator

import (
	"go-diff/service"
	"log"
)

func Validate(args []string) {
	if len(args) < 3 {
		log.Fatal("Usage: go-diff <base_branch> <head_branch>")
	}
	service.CheckGitInstallation()
	service.CheckBranchesExists(args)
}
