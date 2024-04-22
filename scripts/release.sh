#!/usr/bin/bash

confirm_or_die() {
	read -p "Do you want to continue? (y/n)" choice
	case "$choice" in
	y | Y)
		return
		;;
	*)
		exit 2
		;;
	esac
}

# All operations performed relative to the git top level dir
# This means consistent operations regardless of where the script is run from
cd $(git rev-parse --show-toplevel)

NEXT_VERSION="$(git cliff --bumped-version)"

echo "Preparing for a new release [$NEXT_VERSION]"

echo " * Generating candidate CHANGELOG.md"

if [ -f CHANGELOG.staged.md ]; then
	echo "WARNING: Found existing CHANGELOG.staged.md this will be used as is."
	echo "If you want to regenerate a changelog from the latest changes "
	echo "then remove this file before running the script."
	echo
	confirm_or_die
else
	cp CHANGELOG.md CHANGELOG.staged.md
	git cliff -u -t $NEXT_VERSION -p CHANGELOG.staged.md
fi

echo "Tag message"
