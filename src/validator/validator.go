package validator

import (
	"go-diff/helper"
	"go-diff/service"
	"regexp"
)

func Validate(args []string) {
	helper.CheckMatchAndLogError("Usage: go-diff <base_branch> <head_branch> - %v", "Error: missing arguments", len(args) >= 3)
	checkInputArgs(args)

	service.CheckGitInstallation()
	service.CheckBranchesExists(args)
}

func checkInputArgs(args []string) {
	for index, value := range args {
		if index > 0 && index < 3 {
			regexCheck(value)
		}
	}
}

func regexCheck(value string) {
	match, err := regexp.MatchString("^[a-zA-Z0-9_-]+$", value)
	helper.CheckAndLogError("Error checking args: %v", err)
	helper.CheckMatchAndLogError("Invalid branch name: %v", value, match)
}
