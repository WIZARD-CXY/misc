#!/bin/bash

count=10
for ((i=0; i<${count}; i++))
do
{
	echo "${i}-------------------------"
	ssh 223 ls
	echo "${i}*************************"
}&
done
wait

ps