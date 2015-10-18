#include <iostream>
using namespace std;

int lis(int a[],int n){
	int dp[n]={0};
    
    int maxlen=1;

	for(int i=0; i<n; i++){
		dp[i]=1;

		for(int j=0; j<i; j++){
			if(a[j]<=a[i] && dp[j]+1 > dp[i]){
				dp[i]=dp[j]+1;
			}
		}

		maxlen=max(maxlen,dp[i]);
	}

	return maxlen;
}
int main(){
	int d[]={
		5,3,4,8,6,7
	};

	cout<<lis(d,6)<<endl;
}