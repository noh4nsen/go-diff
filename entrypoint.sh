#!/bin/sh

set -eou pipefail

git config --global user.name "go-diff" &&\
git config --global user.email "<>" &&\
git config --global --add safe.directory $GITHHUB_WORKSPACE

cd $GITHUB_WORKSPACE
output=$(go-diff $1 $2) 

echo "full_output='$(echo $output)'" >> $GITHUB_OUTPUT
echo "modified_files='$(jq -n -c --argjson files "$(echo $output | jq -r .modifiedFiles)" '$ARGS.named')'" >> $GITHUB_OUTPUT
echo "modified_dirs='$(jq -n -c --argjson directories "$(echo $output | jq -r .modifiedDirs)" '$ARGS.named')'" >> $GITHUB_OUTPUT
echo "modified_projects='$(jq -n -c --argjson projects "$(echo $output | jq -r .modifiedProjects)" '$ARGS.named')'" >> $GITHUB_OUTPUT