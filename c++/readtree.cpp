#include <vector>
const int max = 1000;
vector<int> G[max];

void readTree(){
        int n;//n node tree, n-1 edges;

        scanf("%d",&n);
        int u,v;
        for(int j=0; j<n-1; j++){
                scanf("%d%d",&u,&v);
                G[u].push_back(v);
                G[v].push_back(u);
        }
}

        
