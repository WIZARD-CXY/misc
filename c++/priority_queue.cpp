#include<queue>
#include<iostream>
#include<vector>
using namespace std;
struct cmp{
	bool operator() (const int a, const int b){
		//return a%10 > b%10; // gewei xiaode priority is higher
        //return a>b; // smaller number priority higher
        return a<b; //bigger number with higer priority,default behaviour
	}
};

int main(){
    vector<int> vec;

    vec.push_back(18);
    vec.push_back(10);
    vec.push_back(11);
    vec.push_back(72);
    vec.push_back(109);
    vec.push_back(14);
    
    priority_queue<int, vector<int> ,cmp> pq;

    for(int i=0; i<vec.size(); i++){
        pq.push(vec[i]);
    }

    while(!pq.empty()){
        int temp=pq.top();
        pq.pop();
        cout<<temp<<endl;
    }
}
