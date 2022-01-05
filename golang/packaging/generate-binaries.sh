#!/bin/bash
## Copyright 2020-2022 Josh Grancell. All rights reserved.
## Use of this source code is governed by an MIT License
## that can be found in the LICENSE file.
BINARY=$1

GOOS=linux GOARCH=amd64 go build -o packaging/artifacts/${BINARY}-linux-amd64
GOOS=darwin GOARCH=amd64 go build -o packaging/artifacts/${BINARY}-darwin-amd64
GOOS=darwin GOARCH=arm64 go build -o packaging/artifacts/${BINARY}-darwin-arm64
GOOS=windows GOARCH=amd64 go build -o packaging/artifacts/${BINARY}-windows-amd64
cd packaging/artifacts
tar -czvf ${BINARY}-linux-amd64.tar.gz ${BINARY}-linux-amd64
tar -czvf ${BINARY}-darwin-amd64.tar.gz ${BINARY}-darwin-amd64
tar -czvf ${BINARY}-darwin-arm64.tar.gz ${BINARY}-darwin-arm64
tar -czvf ${BINARY}-windows-amd64.tar.gz ${BINARY}-windows-amd64