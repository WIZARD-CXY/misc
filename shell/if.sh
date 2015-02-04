#!/bin/bash

#PATH=/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin:~/bin
#export PATH

read -p "Please input (Y/N):" flag

if [ "$flag" == "Y" ] || [ "$flag" == "y" ]; then
    echo "OK,continue!"
    exit 0
elif [ "$flag" == "N" ] || [ "$flag" == "n" ]; then
	echo "Oh,interrupt!"
    exit 0
else
	echo "unsupported command"
    exit 1
fi