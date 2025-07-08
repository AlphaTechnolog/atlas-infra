#!/usr/bin/env bash

set -e

cd "$(dirname "$0")"/../ || exit 1
test -d dist || mkdir -p dist

declare -a LAMBDAS=('atlas-url-listener' 'atlas-url-consumer')

target_lambda="$1"

for lambda in "${LAMBDAS[@]}"; do
  if [[ "$target_lambda" != "" && "$target_lambda" != "$lambda" ]]; then
    continue
  fi

  echo "===Building ${lambda}==="
  GOOS=linux GOARCH=amd64 go build -o bootstrap "./cmd/${lambda}/main.go"
  zip -r "${lambda}.zip" bootstrap
  mv "${lambda}.zip" dist/
  rm bootstrap
done
