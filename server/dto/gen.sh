#!/bin/bash

# iterate go files in current directory
for file in *.go; do
  # extract struct names
  structs=$(grep -E '^type [A-Z][A-Za-z0-9]* struct' "$file" | awk '{print $2}')

  # skip if none
  [ -z "$structs" ] && continue

  # add json tags to the struct
  for s in $structs; do
    gomodifytags \
      -file "$file" \
      -struct "$s" \
      --add-tags json \
      -w
  done
done
