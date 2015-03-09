/*************************************************************************
	> File Name: quicksort.cpp
	> Author: 
	> Mail: 
	> Created Time: Thu 26 Feb 2015 04:48:32 PM CST
 ************************************************************************/

#include<iostream>
using namespace std;

void swap(int &a, int &b){
    int temp=a;
    a=b;
    b=temp;
}

void display(int a[],int n){
    for(int i=0; i<n; i++){
        cout<<a[i]<<" ";
    }
    cout<<endl;
}

int partition(int a[], int p, int r){
    int i=p-1;

    for(int j=p; j<r; j++){
        if(a[j]<a[r]){
            swap(a[++i],a[j]);
        }
    }
    swap(a[++i],a[r]);

    return i;
}
void quicksort(int a[], int p, int r){
    if(p<r){
        int q = partition(a,p,r);
        quicksort(a,p,q-1);
        quicksort(a,q+1,r);
    }
}
int main(){
    int a[100];

    int n;
    while(cin>>n){
        for(int i=0; i<n; i++){
            cin>>a[i];
        }
        quicksort(a,0,n-1);
        display(a,n);
    }
}
