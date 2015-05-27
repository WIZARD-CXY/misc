#include <iostream>
using namespace std;

int lower_bound(int A[], int x, int y, int v){
	while(x<y){
		int m=(x+y)/2;

		if(A[m]>=v){
			y=m;
		}else{
			x=m+1;

		}
	}

	return x;
}

int upper_bound(int A[], int x, int y, int v){
	while(x<y){
		int m=(x+y)/2;

		if(A[m]>v){
			y=m;
		}else{
			x=m+1;
		}
	}
	return x;
}

int bs(int A[], int x, int y, int v){
	while(x<y){
		int m=(x+y)/2;

		if(A[m]==v){
			return m;
		}else if(A[m]>v){
			y=m;
		}else{
			x=m+1;
		}
	}
	return -1;
}

int main(){
	int A[]={-1,1,2,3,3,9,10};

	cout<<lower_bound(A,0,7,4)<<endl;
	cout<<upper_bound(A,0,7,4)<<endl;
	cout<<bs(A,0,7,10)<<endl;

}