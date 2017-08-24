// 0-1 knapsack problem
#include <iostream>
#include <cstdio>
using namespace std;

#define MAXN 1000
#define MAXC 10000

int d[MAXN][MAXC];

int main(){
    int n,C;

    int c,v;

    while(scanf("%d %d", &n, &C) != EOF){

        for (int i=1; i<=n; i++){
            // input every job's cost and value
            scanf("%d %d",&c,&v);
            for(int j=0; j<=C; j++){

                if(i==1){
                    d[i][j]=0;
                }else{
                    d[i][j]=d[i-1][j];
                }
                if (j>=c){
                    d[i][j]=max(d[i][j],d[i-1][j-c]+v);
                }
            }
        }

        printf("max value is %d\n",d[n][C]);
    }
}
