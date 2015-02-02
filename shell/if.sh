#!/bin/bash

#PATH=/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin:~/bin
#export PATH

read -p "Please input (Y/N):" flag

if [ "$flag" == "Y" ] || [ "$flag" == "y" ]; then

   echo "OK,continue!"
   exit 0

fi

if [ "$flag" == "N" ] || [ "$flag" == "n" ]; then

   echo "Oh,interrupt!"
   exit 0

fi

echo "I don't hnow what your choice is" && exit 0