#include <iostream>
#include <algorithm>

using namespace std;

#define MAX 1000

int KK[MAX*MAX];
int target=2000;

bool bs(int x, int n,int A[]){
	int l=0, r=n-1;
	while(l<r){
		int mid=(l+r)/2;

		if(A[mid]==x){
			return true;
		}else if(A[mid]<x){
			l=mid+1;
		}else{
			r=mid;
		}
	}
	return 0;
}

void solve(int A[],int n){

	bool f=false;

	for(int c=0; c<n; c++){
		for(int d=0; d<n; d++){
			KK[c*n+d]=A[c]+A[d];
		}
	}


	sort(KK,KK+n*n);

	bool found=false;

	for(int a=0; a<n; a++){
		for(int b=0; b<n; b++){
			if(bs(target-A[a]-A[b],n*n,KK)){
				found=true;
			}
		}
	}

	if(found){
		cout<<"OK"<<endl;
	}else{
		cout<<"NOT OK"<<endl;
	}
}

int main(){
	int A[]={2,3,4,1,9,89,45,2,4,5,8,9};
	solve(A,12);
}