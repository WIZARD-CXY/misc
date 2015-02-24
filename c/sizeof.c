/*************************************************************************
	> File Name: sizeof.c
	> Author: 
	> Mail: 
	> Created Time: Tue 24 Feb 2015 11:45:01 AM CST
 ************************************************************************/

#include<stdio.h>
typedef struct A{
}AA;

int main(){
    AA a;    
    printf("sizeof empty {} is %d\n",sizeof(a));
}
