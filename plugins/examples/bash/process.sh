#!/bin/bash
#plugin for checking if a process is running

if [[ -n "$1" ]]; then
    if ps aux | grep "$1" > /dev/null; then
        echo Process launch
    else
        echo Process not found
    fi
else
    echo Specify name of process
    echo Ex. /system bash
fi
