#include <iostream>
#include <algorithm>
#include <vector>
#include <cmath>
#include <cstdio>
using namespace std;

int main(){
	freopen("2.txt","r",stdin);

	int n,m;
	scanf("%d%d",&n,&m);

	vector<int> qq(m);
	vector<int> num(n,0);
	// begin[i] record the begin pos of i
	vector<int> begin(n+1);
	int start=0;

	for(int i=0; i<n; i++){
		scanf("%d",&num[i]);
		int count=num[i];

		if(m-start>=count){
			begin[i+1]=start;
			while(count--){
				if(qq[start]!=0){
					//invalidate the old one
					//free the space mark them as 0
					begin[qq[start]]=-1;
					int space=num[qq[start]-1];
                    while(space--){
                    	qq[start+space]=0;
                    }
					
				}
			    qq[(start++)%m]=i+1;
		    }
		}else{
			//set to the head
			start=0;
            begin[i+1]=start;

			while(count--){
				if(qq[start]!=0){
					begin[qq[start]]=-1;
					//invalidate the old one
					//free the space mark them as 0
					int space=num[qq[start]-1];
                    while(space--){
                    	qq[start+space]=0;
                    }
					
				}
			    qq[(start++)%m]=i+1;
		    }
		}
		
	}
   
	for(int i=1; i<=n; i++){
		if(begin[i]!=-1){
			printf("%d %d\n", i, begin[i]);
		}
	}
	
}