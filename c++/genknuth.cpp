void genKnuth(int m, int n){
    // generate m numbers from n records randomly
    for( int i=0; i<n; i++){
        // select m of remaining n-i
        if (bigrand()%(n-i) < m){
            cout<<i<<endl;
            m--;
        }
    }
}

void genset(int m, int n){
     set<int> S;
     while(S.size() < m){
       S.insert(bigrand()%n);
     }


     set<int>::iterator i;
     for( i=S.begin(); i != S.end(); i++){
        cout<<*i<<endl;
     }
}
             
