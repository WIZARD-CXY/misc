#!/bin/bash

data="a  b c d"
for i in ${data#* }
do
	echo $i
done