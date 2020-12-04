#!/bin/bash

RELEASE_VERSION=$(grep -o '[0-9]\.[0-9]\.[0-9]' ./swagger.json)

BUILD_DIR="./dist"

mkdir -p $BUILD_DIR

swagger generate client -f ./swagger.json -A ./basiq-"$RELEASE_VERSION" -t "$BUILD_DIR"
jq 'walk(if type == "object" then with_entries(select(.key | test("x-go-*") | not)) else . end)' swagger.json  > _.json && mv _.json  swagger.json


go test ./test/... -count 1

if [ $? -eq 0 ]; then
  echo "Done generating $RELEASE_VERSION swagger sdk. Check $BUILD_DIR directory."
else
  echo "Failed, please check the error for details."
fi
