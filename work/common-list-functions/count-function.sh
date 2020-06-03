#!/bin/bash

# Usage example:
# ./count-function.sh ": " map fold filter take drop reverse sum zip product maximum minimum

dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

echo "Searching for occurrences in subdirectories of $dir"
for func in "$@"; do
    occurrences="$(rg --type hs " $func" --count-matches | awk -F ':' '{sum += $2} END {print sum}')"
    echo "Found $occurrences occurrences of $func"
done
