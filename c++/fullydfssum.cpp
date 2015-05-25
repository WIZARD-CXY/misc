#include <iostream>
#include <algorithm>
#include <vector>

using namespace std;

int target = 9;
int n=3;

int A[]={2,3,6};
vector<int> path;
vector<vector<int> > paths;

void dfs(int i, int sum){
	if(i==n){
		if(sum==target){         
           paths.push_back(path);
		}
		return;
		
	}

    //not take A[i] into consideration
	dfs(i+1,sum);

    // take A[i] into the consideration
    path.push_back(A[i]);
	dfs(i+1,sum+A[i]);
	path.pop_back();
}

void solve(){
	dfs(0,0);
}

int main(){
	solve();

	for(int i=0; i<paths.size(); i++){
        for(int j=0; j<paths[i].size(); j++){
            cout<<path[j]<<" ";
        }
        cout<<endl;
    }
    
}