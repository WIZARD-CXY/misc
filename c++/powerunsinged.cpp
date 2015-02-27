/*************************************************************************
	> File Name: powerunsinged.cpp
	> Author: 
	> Mail: 
	> Created Time: Fri 27 Feb 2015 12:15:20 PM CST
 ************************************************************************/

#include<iostream>
using namespace std;

double powerunsigned(double base,unsigned int exp){
    if(exp==0){
        return 1;
    }
    if( exp == 1 ){
        return base;
    }

    double res= powerunsigned(base,exp>>1);
    res*=res;

    if(exp&1){
        res*=base;
    }

    return res;
}

int main(){
    cout<<powerunsigned(10.1,9);
}
