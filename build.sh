#!/bin/bash

if [ $# -eq 0 ]; then
  echo "No arguments supplied; add version number e.g ./build 2.0.0"
  exit
fi

# get swagger version file
RELEASE_VERSION=$1

#create dir if not exists
if [ -d "./basiq-swagger-ui" ]; then
  echo "Directory /basiq-swagger-ui exists."
else
  echo "Creating directory"
  cp -r ./swagger-ui ./basiq-swagger-ui
fi

# build a version from a template
cp ./swagger.json ./basiq-swagger-ui/swagger-$RELEASE_VERSION.json

sed -i "s/{{version}}/$RELEASE_VERSION/g" ./basiq-swagger-ui/swagger-$RELEASE_VERSION.json

# generate go client
swagger generate client -f ./swagger.json -A ./basiq-$RELEASE_VERSION
zip -r basiq-sdk-$RELEASE_VERSION.zip ./client ./models
mv basiq-sdk-$RELEASE_VERSION.zip ./basiq-swagger-ui
rm -rf ./client ./models

cp ./logo.svg ./basiq-swagger-ui

cd ./basiq-swagger-ui

# append all swagger versions
i=0
for f in $(ls swagger-*.json | sort -r); do
  version=$(echo $f | grep -o -E '[0-9]+\.[0-9]+\.[0-9]+')
  if [[ $i == 0 ]]; then
    item=$item"{url:"\"$f\"",name:"\"$version\""}"
  else
    item=$item",{url:"\"$f\"",name:"\"$version\""}"
  fi
  i=$(expr $i + 1)
done
url=$item
url="["$url"]"
cd ..
cp ./swagger-ui/index.html ./basiq-swagger-ui
#
cd ./basiq-swagger-ui
sed -i "s/{{swagger-versions}}/$url/g" ./index.html

echo "Done generating $RELEASE_VERSION swagger documentation. Check ./basiq-swagger-ui"
