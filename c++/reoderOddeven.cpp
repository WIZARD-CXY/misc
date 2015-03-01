/*************************************************************************
	> File Name: reoderOddeven.cpp
	> Author: 
	> Mail: 
	> Created Time: Sat 28 Feb 2015 02:25:27 PM CST
 ************************************************************************/

#include<iostream>
using namespace std;

void reorderOddEven(int a[], int length, bool (*cmp)(int)){
    if(a==NULL || length<=0){
        return;

    }

    int begin=0;
    int end=length-1;

    while(begin<end){
        //to find the first even element
        while(begin<end && !cmp(a[begin])){
            begin++;
        }
        //to find the last odd element
        while(begin<end && cmp(a[end])){
            end--;
        }

        int temp=a[begin];
        a[begin]=a[end];
        a[end]=temp;
    }
}

//return true when n is even
bool cmp(int n){
    return (n&1) == 0;
}
bool cmp2(int n){
    return (n%3) != 0;
}
int main(){
    int a[]={1,3,4,5,2,9,10,6,8,7};

    reorderOddEven(a,10,cmp2);

    cout<<a;
}
