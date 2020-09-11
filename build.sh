#!/bin/bash

RELEASE_VERSION=$(grep -o '[0-9]\.[0-9]\.[0-9]' ./swagger.json)

BUILD_DIR="./dist"

mkdir -p $BUILD_DIR

swagger generate client -f ./swagger.json -A ./basiq-"$RELEASE_VERSION" -t "$BUILD_DIR"

go test ./test/...

if [ $? -eq 0 ]; then
  echo "Done generating $RELEASE_VERSION swagger sdk. Check $BUILD_DIR directory."
else
  echo "Failed, please check the error for details."
fi
