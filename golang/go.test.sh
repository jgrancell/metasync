#!/bin/bash

set -e
echo "mode: atomic" > coverage.txt

for FILE in $(go list ./... | grep -v vendor); do
    go test -race -coverprofile=profile.out -covermode=atomic "$FILE"
    if [ -f profile.out ]; then
        sed -i '/^mode/d' profile.out
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done