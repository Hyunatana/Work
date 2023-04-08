#!/bin/bash
PATHS=("/workspaces/codespaces-blank/work/files/*" "/workspaces/codespaces-blank/work/files2/*")
for str in ${PATH[*]}; do
    #FILES=$(find $str ! -newermt "2023-03-29+1 days" -exec realpath {} \;)
    printf "\n$str\n"
done