#!/usr/bin/bash

# All operations performed relative to the git top level dir
# This means consistent operations regardless of where the script is run from
cd $(git rev-parse --show-toplevel)

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
go fmt ./...
go fix ./...

echo
echo "Go dependency update"
go get -u

echo
echo "Go test everything"
go test ./...

echo
echo "Git cliff update the changelog"
git-cliff -l

echo
echo "The next release would be ($(git cliff --bumped-version))"
