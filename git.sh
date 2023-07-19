#!/bin/bash -e
branchName=$(git rev-parse --abbrev-ref HEAD) #Get Current Branch Name
commit_message="$1"
git add .
git commit -m "$commit_message"
branchNum=$(cut -d'-' -f1 <<<"$branchName")
branchAbbr=$(cut -d'-' -f2 <<<"$branchName")
newBranchNum=$((branchNum+1))
git checkout -b "$newBranchNum-$branchAbbr-$(date +%Y.%b.%d.%A_%H-%M-%S)"