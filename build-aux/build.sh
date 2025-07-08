#!/usr/bin/env bash

# This file (re)builds all the layers and lambdas from ts to js and from go to bins
set -e

cd "$(dirname "$0")"/../ || exit 1

# List all sources to build here in a makefile-like sources list.
declare -a JS_SOURCES=(. ./lambdas/js/atlas-url-shortener ./layers/atlas-url-shortener-layer)
declare -a GO_SOURCES=(./lambdas/go:atlas-url-listener)

function xpushd() {
  { pushd "$@" >/dev/null; } 2>&1
}

function xpopd() {
  { popd >/dev/null; } 2>&1
}

function build_js_micros() {
  local sources=${JS_SOURCES[*]}
  for source in ${sources}; do
    xpushd "$source"
      echo "== + [JS] $source + =="
      test -d node_modules || npm install
      npm run build || exit 1
    xpopd
  done
}

function build_go_micros() {
  local sources=${GO_SOURCES[*]}
  for source in ${sources}; do
    local dir="$(echo "$source" | sed 's/:/ /g' | awk '{print $1}')"
    local lambda_name="$(echo "$source" | sed 's/:/ /g' | awk '{print $2}')"
    echo "== + [GO] $lambda_name ($dir) + =="
    xpushd "$dir"
      bash build-aux/build.sh "$lambda_name"
    xpopd
  done
}

build_js_micros
build_go_micros
