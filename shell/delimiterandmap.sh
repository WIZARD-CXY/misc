#!/bin/bash

#data="111,222,333,444,555,666,10.10.103.223,10.10.103.162,10.10.103.224"
#data="111 222 333 10.10.103.223 10.10.103.162 10.10.103.224"
read -p "enter ips > " data
# split data with ','
oldIFS=$IFS  #定义一个变量为默认IFS
#IFS=,        #设置IFS为逗号
# default delimiter is ' '

ii=1
cluster=""
declare -A mm
for i in $data
do
    echo $i
    name="infra"$ii
    item="$name=$i"
    echo $item
    if [ "$ii" == 1 ]; then 
        cluster=$item
    else
        cluster="$cluster,$item"
    fi
    mm[$i]=$name
    let ii++

 done
echo $cluster

for key in ${!mm[@]}; do
    echo "$key value ${mm[$key]} "
done

read -p "enter a ip" IP

echo ${mm[$IP]}
#IFS=$oldIFS  #还原IFS为默认值
