/*************************************************************************
	> File Name: mergeSort.cpp
	> Author: 
	> Mail: 
	> Created Time: Thu 26 Feb 2015 11:18:41 AM CST
 ************************************************************************/

#include<iostream>
using namespace std;

void mergeSort(int a[],int x, int y, int t[]){
    if(y-x>1){
        int m=(x+y)/2;
        int p=x,i=x,q=m;
        mergeSort(a,x,m,t);// recurse on the left side
        mergeSort(a,m,y,t);// recurse on the right side

        while(p<m || q<y){
            if(q>=y || (p<m && a[p]<a[q])){
                t[i++]=a[p++];
            }else{
                t[i++]=a[q++];
            }
        }

        for(int i=x; i<y; i++){
            a[i]=t[i];
        }
    }
}

int main(){
    int a[]={0,-1,4,-9,45,10,23,45};
    int t[10];
    mergeSort(a,0,8,t);

   for(int i=0; i<8; i++){
    cout<<a[i]<<" ";
   }
}
