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
    int n;
    cin>>n;
	cout<<lower_bound(A,0,7,3)<<endl;//prints 3
	cout<<upper_bound(A,0,7,3)<<endl;//prints 5
	cout<<lower_bound(A,0,7,4)<<endl;//prints 5
	cout<<upper_bound(A,0,7,4)<<endl;//prints 5
	cout<<bs(A,0,7,n)<<endl;//prints 6

}
