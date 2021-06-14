
#include<iostream>
using namespace std;
#define OPT(a, b) do{a++; b++;}while(0);

int main(){
    int i=1;
    int j=1;
    if(0)
       OPT(i,j)
    
    cout<<i<<" "<<j<<endl;
}