#!/bin/sh

DARWIN_OS=darwin
LINUX_OS=linux
WIN_OS=windows
ARCH=amd64

BIN_DIR=bin
DARWIN_DIR=$BIN_DIR/$DARWIN_OS/$ARCH
LINUX_DIR=$BIN_DIR/$LINUX_OS/$ARCH
WIN_DIR=$BIN_DIR/$WIN_OS/$ARCH

SRC_FILES=adapters/cli/*.go
BIN_FILE=litepub
VERSION=`git describe --tags --abbrev=0`

echo "Preparing dirs"
rm -rf $BIN_DIR
mkdir -p $DARWIN_DIR
mkdir -p $LINUX_DIR
mkdir -p $WIN_DIR

echo "Building $DARWIN_OS binary"
GOOS=$DARWIN_OS GOARCH=$ARCH go build -o $DARWIN_DIR/$BIN_FILE $SRC_FILES
zip -qj9 $BIN_DIR/$BIN_FILE-$VERSION-$DARWIN_OS-$ARCH.zip $DARWIN_DIR/$BIN_FILE

echo "Building $LINUX_OS binary"
GOOS=$LINUX_OS GOARCH=$ARCH go build -o $LINUX_DIR/$BIN_FILE $SRC_FILES
zip -qj9 $BIN_DIR/$BIN_FILE-$VERSION-$LINUX_OS-$ARCH.zip $LINUX_DIR/$BIN_FILE

echo "Building $WIN_OS binary"
GOOS=$WIN_OS GOARCH=$ARCH go build -o $WIN_DIR/$BIN_FILE.exe $SRC_FILES
zip -qj9 $BIN_DIR/$BIN_FILE-$VERSION-$WIN_OS-$ARCH.zip $WIN_DIR/$BIN_FILE.exe
