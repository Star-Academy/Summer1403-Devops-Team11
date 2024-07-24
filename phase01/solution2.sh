#! /usr/bin/env bash

branches=$(git branch | sed 's/^[* ]//')
IFS=$'\n' read -r -d '' -a branch_array <<< "$branches"
for branch in "${branch_array[@]}";
do
    branch=$(echo "${branch}" | sed 's/^ //')
    git checkout "${branch}"
    find . -type f -exec grep -H "TODO" {} +
done
