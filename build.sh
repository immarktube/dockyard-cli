#!/bin/bash
set -e

APP_NAME=dockyard-cli

platforms=("linux/amd64" "darwin/amd64" "windows/amd64")

for platform in "${platforms[@]}"
do
    GOOS=${platform%%/*}
    GOARCH=${platform##*/}
    output_name="${APP_NAME}-${GOOS}-${GOARCH}"
    [ "$GOOS" == "windows" ] && output_name="${output_name}.exe"

    echo "Building $output_name"
    GOOS=$GOOS GOARCH=$GOARCH go build -o "build/$output_name"
done