#!/bin/sh

set -eoux pipefail

output=$(/go-diff $1 $2) 

modified_files=$(echo $output | jq .modifiedFiles)
modified_dirs=$(echo $outut | jq .modifiedDirs)
modified_projects=$(echo $output | jq .modifiedProjects)

echo "json_output=$output" >> $GITHUB_OUTPUT
echo "modified_files="$modified_files"" >> $GITHUB_OUTPUT
echo "modified_dirs="$modified_dirs"" >> $GITHUB_OUTPUT
echo "modified_projects="$modified_projects"" >> $GITHUB_OUTPUT