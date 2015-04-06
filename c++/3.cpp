#include <iostream>
#include <cmath>
#include <cstdio>
using namespace std;

#define MAXN 20000
int w[MAXN][MAXN];
bool vis[MAXN];
int dist[MAXN];
int x[MAXN];
int y[MAXN];

int min(int a,int b){
	if(a<b){
		return a;
	}else {
		return b;
	}
}

void dijk(int n, int src){

      for(int i=0; i<n; i++){
          dist[i]=w[src][i];
      }
      
      dist[src]=0;
      vis[src]=1;

      for(int j=0; j<n; j++){

           int minDst=0x7fffffff;
           int minIdx=-1;
           for(int i=0; i<n; i++){
                if(!vis[i] && dist[i] < minDst){
                    minDst=dist[i];
                    minIdx= i;
                }
            }

	       vis[minIdx]=1;

	        for(int i=0; i<n; i++){
	            if(!vis[i] && w[minIdx][i]<0x7fffffff && dist[i] >=(dist[minIdx]+w[minIdx][i])){
	                    dist[i]=dist[minIdx]+w[minIdx][i];
	              
	            }
	        }
    }
}  

int main(){
	freopen("3.txt","r",stdin);

	int n;
	cin>>n;

	for(int i=0; i<n; i++){
		cin>>x[i]>>y[i];
	}

	for(int i=0; i<n; i++){
		for(int j=i+1; j<n; j++){
			w[i][j]=w[j][i]=min(abs(x[i]-x[j]),abs(y[i]-y[j]));
		}
	}

	dijk(n,0);
	cout<<dist[n-1]<<endl;
	
}