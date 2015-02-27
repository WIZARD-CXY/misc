/*************************************************************************
	> File Name: numberof1.cpp
	> Author: 
	> Mail: 
	> Created Time: Fri 27 Feb 2015 10:47:53 AM CST
 ************************************************************************/

#include<iostream>
using namespace std;


int numberof1(int n){
    int count=0;

    while(n){
        count++;
        n=n&(n-1);
    }
    return count;
}

int main(){
    cout<<numberof1(8)<<endl;
}
