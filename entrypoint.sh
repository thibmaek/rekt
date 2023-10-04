#! /usr/bin/env bash

set -e

if [ -z "$1" ]; then
  echo "No apk file supplied"
  exit 1
fi

rekt intro

if [ -d "./scan/app" ]; then
  echo "Cleaning previous decompiled app at ./scan/app ..."
  rm -rf ./scan/app
fi

rekt decompile -apk="$1" -outputDir="./scan/app"
rekt probe -inputDir="./scan/app"
rekt break -inputDir="./scan/app"
