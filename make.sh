#!/bin/bash

set -e

PROJECT_NAME="tor_shinanigans"
BUILD_DIR="bin"   # <- You had a typo here: "BUIDL_DIR"
PIC_DIR="pictures"
case "$1" in
    build)
        echo "Building..."
        mkdir -p "$BUILD_DIR"          # Make sure the bin/ directory exists
        go build -o "$BUILD_DIR/$PROJECT_NAME" main.go
        ;;
    run)
        echo "Running..."
        go run main.go
        ;;
    clean)
        echo "Cleaning..."
        rm -rf "$BUILD_DIR"             # <- You had $Building, should be $BUILD_DIR
        rm -rf "$PIC_DIR"             # <- You had $Building, should be $BUILD_DIR
        ;;
    *)
        echo "Usage: $0 {build|run|clean}"
        exit 1
        ;;
esac

