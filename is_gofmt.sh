#!/bin/bash

set -eu

if [ -n "$(gofmt -l .)" ]; then
  echo "Go code is not formatted:"
  gofmt -d .
  exit 1
fi
