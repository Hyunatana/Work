#!/bin/bash
array=($(df -h | awk '{print $5}' | sed 's/%//g' | sed '1d'))
FPATH=("/workspaces/codespaces-blank/testing/files/*" "/workspaces/codespaces-blank/testing/files2/*")
for i in ${array[@]}; do
    if [ $i -ge 40 ]; then
        printf "\nDisk usage is greater than 80%% removing old files\n"
        for i in ${FPATH[@]}; do
            FILES=$(find $i -mtime +21 -exec realpath {} \; 2>/dev/null)
            rm -rf $FILES
        done
        printf  "\nRemoving old files from $FPATH\n"
    fi
done