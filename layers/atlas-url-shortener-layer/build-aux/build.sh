#!/usr/bin/env bash

cd "$(dirname "$0")"/../ || exit 0

test -d dist && rm -rf dist
npx tsc
mkdir -pv out/nodejs
cp -rvf dist/* out/nodejs
cp -rvf node_modules out/nodejs
rm -rf dist && mv -v out dist
cd dist && zip -r ../dist/layer.zip . && rm -rf nodejs && cd ..
