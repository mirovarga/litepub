#!/bin/sh

LINUX_OS=linux
WIN_OS=windows
DARWIN_OS=darwin
ARCH=amd64

BIN_FILE=litepub
BIN_DIR=bin
LINUX_DIR=$BIN_DIR/$LINUX_OS/$ARCH
WIN_DIR=$BIN_DIR/$WIN_OS/$ARCH
DARWIN_DIR=$BIN_DIR/$DARWIN_OS/$ARCH

OTHER_FILES="LICENSE README.md"
VERSION=`git describe --tags --abbrev=0`

echo "Preparing dirs"
rm -rf $BIN_DIR
mkdir -p $LINUX_DIR
mkdir -p $WIN_DIR
mkdir -p $DARWIN_DIR

echo "Building $LINUX_OS distribution"
GOOS=$LINUX_OS GOARCH=$ARCH go build -o $LINUX_DIR/$BIN_FILE
zip -qj9 $BIN_DIR/$BIN_FILE-$VERSION-$LINUX_OS-$ARCH.zip $LINUX_DIR/$BIN_FILE $OTHER_FILES

echo "Building $WIN_OS distribution"
GOOS=$WIN_OS GOARCH=$ARCH go build -o $WIN_DIR/$BIN_FILE.exe
zip -qj9 $BIN_DIR/$BIN_FILE-$VERSION-$WIN_OS-$ARCH.zip $WIN_DIR/$BIN_FILE.exe $OTHER_FILES

echo "Building $DARWIN_OS distribution"
GOOS=$DARWIN_OS GOARCH=$ARCH go build -o $DARWIN_DIR/$BIN_FILE
zip -qj9 $BIN_DIR/$BIN_FILE-$VERSION-$DARWIN_OS-$ARCH.zip $DARWIN_DIR/$BIN_FILE $OTHER_FILES
