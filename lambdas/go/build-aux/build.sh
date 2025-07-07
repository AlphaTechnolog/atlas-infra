#!/usr/bin/env bash

set -e

cd "$(dirname "$0")"/../ || exit 1
test -d dist || mkdir -p dist

declare -a LAMBDAS=('atlas-url-listener')

for lambda in "${LAMBDAS[@]}"; do
  echo "===Building ${lambda}==="
  GOOS=linux GOARCH=amd64 go build -o bootstrap "./cmd/${lambda}/main.go"
  zip -r "${lambda}.zip" bootstrap
  mv "${lambda}.zip" dist/
  rm bootstrap
done
