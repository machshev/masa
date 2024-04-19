#!/usr/bin/bash

adrs generate toc >doc/adr/toc.md

go mod tidy
go get -u
go test ./...

git-cliff -o
