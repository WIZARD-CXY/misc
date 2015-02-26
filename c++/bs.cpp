/*************************************************************************
	> File Name: bs.cpp
	> Author: 
	> Mail: 
	> Created Time: Thu 26 Feb 2015 11:15:21 AM CST
 ************************************************************************/

#include<iostream>
using namespace std;


int bs(int a[],int l, int r, int v){
    while(l<r){
        int m=(l+r)/2;

        if(a[m]==v){
            return m;
        }else if(a[m]>v){
            r=m;
        }else{
            l=m+1;
        }
    }
    return -1;
}

int main(){
    int a[]={0,1,2,3,4,5,6,7,8,9};

    cout<<bs(a,0,10,9)<<endl;
}
