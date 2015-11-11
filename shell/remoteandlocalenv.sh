// "" will make variable expansion locally,
// '' will make variable remotely
ssh user@ip echo 'PWD is ${PWD}' //remote PWD
ssh user@ip echo "PWD is ${PWD}" // local PWD
