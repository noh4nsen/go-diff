#!/bin/sh

output=$(./go-diff $1 $2)
echo $output
echo "modified_output=output" >> $GITHUB_OUTPUT