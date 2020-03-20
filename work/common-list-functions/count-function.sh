#!/bin/bash

dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

echo "Searching for occurrences in subdirectories of $dir"
for func in "$@"; do
    occurrences="$(grep --include \*.hs -Ro " $func" "$dir" | wc -l | xargs)"
    echo "Found $occurrences occurrences of $func"
done
