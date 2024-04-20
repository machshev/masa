#!/usr/bin/bash

# This script is used to tidy everything and prepare for a new release
echo
echo "Update ADR doc table of contents"
adrs generate toc >doc/adr/toc.md

echo
echo "regenerate protobuf code"
(
	cd proto
	buf generate
)

echo
echo "Go tidy"
go mod tidy

echo
echo "Go dependency update"
go get -u

echo
echo "Go test everything"
go test ./...

echo
echo "Git cliff update the changelog"
git-cliff -o

echo
