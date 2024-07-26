#!/usr/bin/env bash 


git for-each-ref --format '%(refname:short)' refs/heads/ | while read branch; do
    echo "${branch}:"
    git checkout "${branch}"
    find . -type f -exec grep -H "TODO" {} +
done
