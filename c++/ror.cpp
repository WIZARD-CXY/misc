/*************************************************************************
	> File Name: ror.cpp
	> Author: 
	> Mail: 
	> Created Time: Mon 09 Mar 2015 11:47:52 PM CST
 ************************************************************************/

#include<iostream>
using namespace std;


int gcd(int a, int b){
    return b==0?a:gcd(b,a%b);
}

void ror(int array[], int n, int k){
    int i;
    const int g=gcd(n,k);

    k%=n;

    for(i=0; i<g; i++){
        int j=i;

        int cur=array[j],tmp;

        do{
            tmp=array[(j+k)%n];
            array[(j+k)%n]=cur;
            cur=tmp;
            j=(j+k)%n;
            cout<<"j "<<j<<" ";
        }while(j!=i);
        cout<<endl;
    }
}

int main(){
    int i;
    int a[]={1,2,3,4,5,6};

    ror(a,6,3);

    for(i=0; i<6; i++){
        cout<<a[i]<<" ";
    }
    cout<<endl;
}
