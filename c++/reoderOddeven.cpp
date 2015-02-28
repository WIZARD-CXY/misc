/*************************************************************************
	> File Name: reoderOddeven.cpp
	> Author: 
	> Mail: 
	> Created Time: Sat 28 Feb 2015 02:25:27 PM CST
 ************************************************************************/

#include<iostream>
using namespace std;

void reorderOddEven(int a[],int length){
    if(a==NULL || length<=0){
        return;

    }

    int begin=0;
    int end=length-1;

    while(begin<end){
        //to find the first even element
        while(begin<end && (a[begin]&0x01) !=0){
            begin++;
        }
        //to find the last odd element
        while(begin<end && (a[end]&0x01) ==0){
            end--;
        }

        int temp=a[begin];
        a[begin]=a[end];
        a[end]=temp;
    }
}

int main(){
    int a[]={1,3,4,5,2,9,10,6,8,7};

    reorderOddEven(a,10);

    cout<<a;
}
