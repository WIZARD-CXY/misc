#include <iostream>
#include <cstdio>
using namespace std;

int main(){
	int x=10039;
	printf("%x\n",x);
    
    //odd bit right shift , even bits left shift then do bitwise or
	x=((x&0xaaaa)>>1) | ((x&0x5555)<<1);
    printf("%x\n",x);
}