#! /bin/bash

branches=$(git branch | sed 's/^[* ]//')

IFS=$'\n' read -r -d '' -a branch_array <<< "$branches"
for branch in "${branch_array[@]}";
do
	echo "$branch"
done
