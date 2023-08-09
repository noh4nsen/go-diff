package service

import (
	"go-diff/core"
	"go-diff/helper"
	"go-diff/model"
	"log"
	"os/exec"
)

func ExecDiff(args []string) model.Changes {
	execFetch()

	cmd := exec.Command("git", "diff", "--name-only", args[1], args[2])

	output, err := cmd.Output()
	helper.CheckAndLogError("Error executing 'git diff' command: %v", err)
	log.Println("Executed: 'git diff --name-only", args[1], args[2], "'")

	return core.ExtractChanges(string(output))
}

func CheckGitInstallation() {
	gitPath, err := exec.LookPath("git")
	helper.CheckAndLogError("Git is not installed on the system:%v", err)
	log.Println("Git is installed: ", gitPath)
}

func CheckBranches(args []string) {
	for index, branch := range args {
		if index > 0 && index < 3 {
			execCheckout(branch)
			execPull(branch)
			execRevParse(branch)
		}
	}
}

func execFetch() {
	cmd := exec.Command("git", "fetch")
	_, err := cmd.Output()
	helper.CheckAndLogError("Error executing 'git fetch' command: %v", err)
	log.Println("Executed: 'git fetch'")
}

func execCheckout(branch string) {
	cmd := exec.Command("git", "checkout", branch)

	_, err := cmd.Output()
	helper.CheckAndLogError("Error checking out branch: %v", err)
	log.Println("Checkout branch: ", branch)
}

func execPull(branch string) {
	cmd := exec.Command("git", "pull")

	_, err := cmd.Output()
	helper.CheckAndLogError("Error pulling branch: %v", err)
	log.Println("Pull branch: ", branch)
}

func execRevParse(branch string) {
	cmd := exec.Command("git", "rev-parse", "--verify", branch)

	output, err := cmd.Output()
	helper.CheckAndLogError("Error verifying branch: %v", err)
	log.Println("Comparing on hash: ", string(output))
}
