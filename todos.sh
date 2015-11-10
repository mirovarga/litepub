#!/bin/sh

grep -IRn --exclude=`basename "$0"` --exclude-dir=.git --color=auto 'TODO\|FIXME\|XXX' .
