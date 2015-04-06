#include <iostream>
#include <algorithm>
#include <vector>
#include <cmath>
#include <cstdio>
using namespace std;

int main(){
	freopen("1.txt","r",stdin);
	int x,y,z;
	cin>>x>>y>>z;

	string input;
	cin>>input;
	//count[0] means R, 1 means Y ,2 means B
	int count[3]={0};
	int max=0;
	int current=0;

	int len=input.size();
	int i;
	bool varnish=0;

	for(i=0; i<len; i++){

		if(input[i]=='R'){
			count[0]++;
		}else if(input[i]=='Y'){
			count[1]++;
		}else{
			count[2]++;
		}
		current++;

		if((abs(count[0]-count[1])==x && 
		   abs(count[1]-count[2])==y &&
		   abs(count[2]-count[0])==z) || 
			(abs(count[0]-count[1])==y && 
		   abs(count[1]-count[2])==x &&
		   abs(count[2]-count[0])==z) ||
            (abs(count[0]-count[1])==y && 
		   abs(count[1]-count[2])==z &&
		   abs(count[2]-count[0])==x) ||
            (abs(count[0]-count[1])==z && 
		   abs(count[1]-count[2])==x &&
		   abs(count[2]-count[0])==y) ||
            (abs(count[0]-count[1])==z && 
		   abs(count[1]-count[2])==y &&
		   abs(count[2]-count[0])==x) ||
            (abs(count[0]-count[1])==x && 
		   abs(count[1]-count[2])==z &&
		   abs(count[2]-count[0])==y)){
			varnish=1;
		}

		if(varnish){
			if(current>max){
				max=current;
			}
			varnish=0;
			current=0;
			count[0]=count[1]=count[2]=0;
		}else{
			if(current>max){
				max=current;
			}

		}
           
	}
	cout<<max<<endl;
}