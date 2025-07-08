#!/usr/bin/env bash

set -e

cd "$(dirname "$0")"/../ || exit 1

test -d dist && rm -rf dist
npx tsc
cd dist/ || exit 1
rm -rf layers
mv lambdas/js/atlas-url-visit-count-subscription-manager/src/* ./ && rm -rf lambdas
