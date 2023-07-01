#!/usr/bin/env bash

download-package() {
  set -eo pipefail
  local PACKAGE_NAME=$1
  local PACKAGE_VERSION=$2
  local PACKAGE_CHECKSUM=$3
  local PACKAGE_FILE="$PACKAGE_NAME-$PACKAGE_VERSION.tgz"

  # download the tarball from NPM's registry
  mkdir -p "node_modules/$PACKAGE_NAME"
  curl "https://registry.npmjs.org/$PACKAGE_NAME/-/$PACKAGE_FILE" >"node_modules/$PACKAGE_FILE"

  # verify the tarball's SHA512 checksum
  local PACKAGE_SHA
  PACKAGE_SHA=$(shasum -b -a 512 "node_modules/$PACKAGE_FILE" | awk '{ print $1 }' | xxd -r -p | base64 -w0)
  [ "$PACKAGE_SHA" == "$PACKAGE_CHECKSUM" ]

  # uncompress the package into a directory matching its name
  gunzip -dc "node_modules/$PACKAGE_FILE" | tar -xf - --strip-components 1 -C "node_modules/$PACKAGE_NAME"
}

# install by manually providing metadata from https://registry.npmjs.org/solid-js
download-package solid-js 1.6.6 "5x33mEbPI8QLuywvFjQP4krjWDr8xiYFgZx9KCBH7b0ZzypQCHaUubob7bK6i+1u6nhaAqhWtvXS587Kb8DShA=="
