#include <iostream>
#include <algorithm>
#include <vector>
#include <cmath>
#include <cstdio>
using namespace std;

int main(){
	freopen("1.txt","r",stdin);
	string line;
    string res="";

    while(getline(cin,line)){
    	if(line.size()==0){
    		continue;
    	}
    	bool isfisrtspace=true;
    	bool isCap=true;
    	bool isleadingspace=true;
        string tmp;

		for(int i=0; i<line.size(); i++){
            //transeform to capital
            while(isleadingspace && line[i]==' '){
            	i++;
            }
            isleadingspace=false;
			if(isCap){
				if('a'<=line[i] && line[i]<='z'){
					tmp.push_back(line[i]-'a'+'A');
					isCap=false;
					continue;
				}else if (line[i]>='A' && line[i] <='Z'){
					tmp.push_back(line[i]);
					isCap=false;
					continue;
				}
				
			}

			if(isfisrtspace && line[i]==' '){
				tmp.push_back(' ');
				isfisrtspace=false;
			}

			if(line[i]==',' || line[i]=='.'){
				tmp.push_back(line[i]);

				//remove ' ' before ','
				string::iterator it=tmp.end()-2;
				while(it!=tmp.begin() && *it==' '){
					tmp.erase(it);
				}
				isfisrtspace=true;

				if(line[i]=='.'){
					//to the end
					isCap=true;
				}
			}
			if(!isCap){
				if(line[i]>='a' && line[i] <='z'){
					tmp.push_back(line[i]);
					isfisrtspace=true;
				}else if (line[i]>='A' && line[i] <='Z'){
					tmp.push_back(line[i]-'A'+'a');
					isfisrtspace=true;
				}
				
			}

		}
		string::iterator it=tmp.end()-1;
		// just ' '
		if(it==tmp.begin()){
			continue;
		}
		while(it!=tmp.begin() && *it==' '){
			tmp.erase(it);
		}

		res+=(tmp+'\n');
    }

    cout<<res;
}