#! /bin/bash

git for-each-ref --format '%(refname:short)' refs/heads/ | xargs -I {} sh -c "git checkout {} && echo {}':\n'&& grep "TODO" *"


#another solution:

#branches=$(git branch | sed 's/^[* ]//')
#IFS=$'\n' read -r -d '' -a branch_array <<< "$branches"
#for branch in "${branch_array[@]}";
#do
#	echo "$branch :\n"
#	...	
#done
