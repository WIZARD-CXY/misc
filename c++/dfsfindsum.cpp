#include <iostream>
#include <algorithm>
#include <vector>

using namespace std;

int target = 5;
int n=8;

int A[]={2,3,4,1,67,34,-20,45};
vector<int> path;

bool dfs(int i, int sum){
	if(i==n){
		if(sum==target){         
           return true;
		}
		return false;
	}

    //not take A[i] into consideration
	if(dfs(i+1,sum)) return true;

    // take A[i] into the consideration
	if(dfs(i+1,sum+A[i])) {
		path.push_back(A[i]);
		return true;
	}

	return false;
}

void solve(){
	if(dfs(0,0)){
		cout<<"YES"<<endl;
	}else{
		cout<<"NO"<<endl;
	}
}

int main(){
	
	solve();

	for(int i=0; i<path.size(); i++){
        cout<<path[i]<<" ";
    }
    cout<<endl;
}