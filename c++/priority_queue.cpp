#include<priority_queue>

struct cmp{
	bool operator() (const int a, const int b){
		return a%10 > b%10; 
	}
};

priority_queue<int, vector<int>, cmp> q; // gewwei digit bigger priority is smaller