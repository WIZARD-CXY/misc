/*************************************************************************
	> File Name: print1toNdigits.cpp
	> Author: 
	> Mail: 
	> Created Time: Fri 27 Feb 2015 03:47:04 PM CST
 ************************************************************************/

#include<iostream>
#include<cstring>

using namespace std;

void PrintNum(char *num){
    int length = strlen(num);
    int i=0;
    while(num[i]=='0' && i<length){
        i++;
    }

    for(; i<length; i++){
        cout<<num[i];
    }
    cout<<endl;
}


void generate(char num[], int n, int index){

    if(index == n){
        PrintNum(num);
        return;

    }
    for(int i=0; i<10; i++){
        num[index]=i+'0';
        generate(num,n,index+1);
    }
}

void Print1toNdigits(int n){
    if(n<=0){
        return; 
    }

    char number[n+1];
    number[n]='\0';


    generate(number,n,0);
}
int main(){
    Print1toNdigits(9);
}
