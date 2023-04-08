#!/bin/bash
DATE=$(date +%F)
LOG=/workspaces/codespaces-blank/work/$DATE-cdt-nas-purge.log
ARRAY=("/workspaces/codespaces-blank/testing/files/*" "/workspaces/codespaces-blank/testing/files2/*")
# Make log file if it doesnt exist
if [[ ! -f $LOG ]]; then
    mkdir -p /workspaces/codespaces-blank/work/files >> /dev/null
    touch $LOG
fi

printf "$(date)" > $LOG

# Delete files based on user inputed date.

while true; do
read -p "Enter the date you would like to begin removing files from (YYYY-MM-DD): " INPUT
    if [[ $INPUT =~ ^[0-9]{4}-[0-9]{2}-[0-9]{2}$ && "$INPUT" -ge "$DATE" ]]; then
    read -p "This is not reversible. Are you sure?  [Y/n] " -n 1 -r
        if [[ $REPLY =~ ^[Yy]$ ]]; then
            for i in ${ARRAY[@]}; do
                printf "\nRemoving files from beginning at $INPUT\n" >> $LOG
                FILES=$(find $i ! -newermt "$INPUT+1 days" -exec realpath {} \; 2>/dev/null)
                if [[ -z $FILES ]]; then
                    printf "No files to remove from on or before the selected date.\n" >>  $LOG
                else
                    printf "Files removed:\n$FILES\n" >> $LOG
                    rm -rf $FILES
                fi
            done
            exit 0
        fi
    else
        printf "Invalid input, please enter a valid date (YYYY-MM-DD). Please try again.\n"
    fi
done