/*************************************************************************
	> File Name: quicksort.cpp
	> Author: 
	> Mail: 
	> Created Time: Thu 26 Feb 2015 04:48:32 PM CST
 ************************************************************************/

#include<iostream>
using namespace std;

void swap(int &a, int &b){
    // avoid a==b situation
    if(a==b){
        return;
    }
    a=a^b;
    b=b^a;
    a=a^b;
}

void display(int a[],int n){
    for(int i=0; i<n; i++){
        cout<<a[i]<<" ";
    }
    cout<<endl;
}

int partition(int a[], int l, int r){
    int i=l-1;

    for(int j=l; j<r; j++){
        //use a[r] as pivot
        if(a[j]<a[r]){
            swap(a[++i],a[j]);
        }
    }
    swap(a[++i],a[r]);
    return i;
}
void quicksort(int a[], int l, int r){
    if(l<r){
        int p=partition(a,l,r);
        quicksort(a,l,p-1);
        quicksort(a,p+1,r);
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
