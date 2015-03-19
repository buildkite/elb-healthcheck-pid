#!/bin/bash
set -e

NAME="elb-healthcheck-pid"
BUILD_PATH="tmp/build"

function build {
  echo "--- Building binary for $1/$2"

  export GOOS=${1}
  export GOARCH=${2}

  echo "GOOS=$GOOS"
  echo "GOARCH=$GOARCH"

  BINARY_NAME="$NAME-$1-$2"

  go build -v -o $BUILD_PATH/$BINARY_NAME *.go
  chmod +x $BUILD_PATH/$BINARY_NAME

  echo -e "\nDone: \033[33m$BUILD_PATH/$BINARY_NAME\033[0m 💪"
}

rm -rf $BUILD_PATH
mkdir -p $BUILD_PATH

build "linux" "amd64"
build "linux" "386"
build "darwin" "386"
build "darwin" "amd64"
