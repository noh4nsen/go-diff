#!/bin/sh

output=$(/go-diff $1 $2) 
modified_files=$(echo $output | jq .modifiedFiles)
modified_dirs=$(echo $output | jq .modifiedDirs)
modified_projects=$(echo $output | jq .modifiedProjects)

echo "modified_files="$modified_files"" >> $GITHUB_OUTPUT
echo "modified_dirs="$modified_dirs"" >> $GITHUB_OUTPUT
echo "modified_projects="$modified_projects"" >> $GITHUB_OUTPUT