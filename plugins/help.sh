#!/bin/bash
#plugin for display available plugins

cd plugins

for file in *.sh; do
    if [[ -f "$file" ]]; then
        first_line=$(head -n 1 "$file")
        second_line=$(head -n 2 "$file" | tail -n 1)
        if [[ "$second_line" == \#* ]]; then
            filename="${file%.sh}"
            echo "/$filename - ${second_line:1}"
        fi
    fi
done

echo "Created by Konare1ka"