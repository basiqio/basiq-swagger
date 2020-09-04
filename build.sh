#!/bin/bash

RELEASE_VERSION=$(grep -o '[0-9]\.[0-9]\.[0-9]' ./swagger.json)

BUILD_DIR="./dist"

if [ -d "$BUILD_DIR" ]; then
  echo "Directory  exists."
else
  echo "Creating directory"
  mkdir "$BUILD_DIR"
fi

swagger generate client -f ./swagger.json -A ./basiq-"$RELEASE_VERSION"

mv ./client "$BUILD_DIR"
mv ./models "$BUILD_DIR"
rm -rf ./client ./models

echo "Done generating $RELEASE_VERSION swagger sdk. Check "$BUILD_DIR
