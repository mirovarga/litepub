#!/bin/sh

OPTIONS="-RnI --exclude=`basename "$0"` --exclude-dir=.git --color=auto"

grep $OPTIONS 'TODO' .
grep $OPTIONS 'FIXME' .
grep $OPTIONS 'XXX' .
