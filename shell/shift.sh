#!/bin/bash
# shift command to shift leftwise

set -e
echo "----"
echo $OPTIND

echo "----"

echo $#
echo $@

shift 2

echo "----"
echo $OPTIND

echo "----"

echo $#
echo $1

shift 

echo $#
echo $3


