#include <iostream>
#include <algorithm>
#include <vector>

using namespace std;

int target = -1;
int n=6;

// must add one dummy element in the end in case dfs search search A[n]
int A[]={2,3,6,5,4,-1};
vector<vector<int> > paths;

void dfs(int i, int sum, vector<int> &path){
	if(i==n){
		if(sum==target){         
           paths.push_back(path);
		}
		return;
	}

    //not take A[i] into consideration
	dfs(i+1,sum,path);

    // take A[i] into the consideration
    path.push_back(A[i]);
	dfs(i+1,sum+A[i],path);
	path.pop_back();
}

int main(){
	vector<int> path;
	dfs(0,0,path);

	for(int i=0; i<paths.size(); i++){
        for(int j=0; j<paths[i].size(); j++){
            cout<<paths[i][j]<<" ";
        }
        cout<<endl;
    }
    
}