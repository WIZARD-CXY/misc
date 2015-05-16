#!/bin/bash
# This shell script is used to retag the image

set -e

if [ -z "$1" ]; then
	echo "must use new reg as first param"
	exit 1
fi

imageNum=$(docker images | wc | awk '{print $1}')
((imageNum--))
imagesIDs=$(docker images | awk '{print $3}')
read -a imageIDarray <<< ${imagesIDs}
imageRepos=$(docker images | awk '{print $1}')
read -a imageRepoarray <<< $imageRepos
imagetags=$(docker images | awk '{print $2}')
read -a imagetagarray <<< $imagetags

# $1 is the new registry name
for ((i=1; i<=${imageNum}; i++)); do
	#echo "${imageIDarray[$i]}"
	docker tag -f "${imageIDarray[${i}]}" "${1}"/"${imageRepoarray[${i}]#*/}":"${imagetagarray[${i}]}"
done