#include <iostream>
#include <cmath>
#include <vector>
using namespace std;

int main(){
	int n;
	cin>>n;

	vector<int> vis(n/32+1,0);

	vector<int> primes;

	for(int i=2; i<n;i++){
		if((vis[i/32] & (1<<(i%32)))==0){
			primes.push_back(i);

			for(int j=i; j<n; j+=i){
				vis[j/32] |= (1<<(j%32));
			}
		}
	}

	for(auto &p: primes){
		cout<<p<<" ";
	}
}