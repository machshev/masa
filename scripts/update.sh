#!/usr/bin/bash

# All operations performed relative to the git top level dir
# This means consistent operations regardless of where the script is run from
cd "$(git rev-parse --show-toplevel)" || exit

# This script is used to tidy everything and prepare for a new release
echo " * Update ADR doc table of contents"
adrs generate toc >doc/adr/toc.md

echo " * regenerate protobuf code"
(
	cd proto || exit 10
	buf generate
)

echo " * Go tidy"
go mod tidy
go fmt ./...
go fix ./...

echo " * Go dependency update"
go get -u ./cmd/masa
go get -u ./cmd/masa-server

echo " * Go Vet"
if ! go vet ./...; then
	echo "***** Aborting (Vet failed) *****"
	exit 1
fi

echo " * Build Masa Server"
if ! go build -o bin/masa-server ./cmd/masa-server; then
	echo "***** Aborting (build failed) *****"
	exit 2
fi

echo " * Build Masa CLI"
if ! go build -o bin/masa ./cmd/masa; then
	echo "***** Aborting (build failed) *****"
	exit 3
fi

echo " * Go test everything"
if ! go test -shuffle on ./...; then
	echo "***** Aborting *****"
	exit 4
fi

echo " * Git cliff update the changelog"
git-cliff -u

echo
echo "[The next release would be $(git cliff --bumped-version)]"
