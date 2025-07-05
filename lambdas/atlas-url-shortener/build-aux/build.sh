#!/usr/bin/env bash

cd "$(dirname "$0")"/../ || exit 1

test -d dist && rm -rf dist
npx tsc
cd dist || exit 1
rm -rf atlas-url-shortener-layer && mv atlas-url-shortener/src/* ./
for x in atlas-url-shortener/{src,}; do rmdir "$x"; done