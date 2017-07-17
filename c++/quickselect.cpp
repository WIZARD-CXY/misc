#include <iostream>
#include <vector>
#include <cstdlib>
#include <ctime>
using namespace std;

void swap(int &a, int &b){
    int t=a;
    a=b;
    b=t;
}
// randint returns a random integer from [l, r]
int randint(int l, int r){
    return l+(rand()%(r-l+1));
}
void print1(int *a){
    for (int i=0; i<5; i++){
        printf("-- %d ",a[i]);
    }
    printf("\n");
}

void print2(int *a){
    for (int i=0; i<5; i++){
        printf("xx %d ",a[i]);
    }
    printf("\n");
}

int select(int *a, int l, int u, int k){
    if(l<u){
        int m=l;
        // use random pivot        
        swap(a[l], a[randint(l,u)]);
        
        for(int i=l+1; i<=u; i++){
            if(a[i]<a[l]){
                swap(a[++m], a[i]);
            }
            //print1(a);
        }
        swap(a[l], a[m]);
        //print2(a);

        if (k<m){
            // recurse on the first half
            return select(a, l, m-1, k);
        }else if (k>m){
            // recurse on the second half
            return select(a, m+1, u, k);
        }else {
            return a[m];
        }
    }
    return a[l];
}

// loop select use loop to find k-th number instead of recursion
int loopselect(int *a, int l, int u, int k){
    while(l<u){
        int m=l;
        swap(a[l],a[randint(l,u)]);

        for(int i=l+1; i<=u; i++){
            if (a[i]<a[l]){
                swap(a[++m],a[i]);
            }
        }

        swap(a[m],a[l]);

        if(m<k){
            l=m+1;
        }else if(m>k){
            u=m-1;
        }else{
            return a[m];
        }
    }
    return a[l];
}



int main(){
    srand(time(0));

    int a[6]={1,4,0,3,4,10};
    int p = loopselect(a,0,5,5);
    printf("%d\n",p);
}