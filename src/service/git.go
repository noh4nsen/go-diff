package service

import (
	"go-diff/core"
	"go-diff/helper"
	"go-diff/model"
	"log"
	"os/exec"
)

func ExecDiff(args []string) model.Changes {
	cmd := exec.Command("git", "diff", "--name-only", args[1], args[2])

	output, err := cmd.Output()
	helper.CheckAndLogError("Error executing 'git diff' command: %v", err)

	return core.ExtractChanges(string(output))
}

func CheckGitInstallation() {
	gitPath, err := exec.LookPath("git")

	helper.CheckAndLogError("Git is not installed on the system:%v", err)
	log.Println("Git is installed: ", gitPath)
}

func CheckBranchesExists(args []string) {
	for index, value := range args {
		if index != 0 {
			cmd := exec.Command("git", "rev-parse", "--verify", value)

			output, err := cmd.Output()
			helper.CheckAndLogError("Error verifying branch: %v", err)
			log.Println("Branch: ", string(output))
		}
	}
}
