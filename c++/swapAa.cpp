#include <iostream>
#include <cstring>
using namespace std;

/*void swap(char *a, char *b){
	int temp=*a;
	*a=*b;
	*b=temp;
}*/
void proc1(string &str){
	int len=str.size();

    int i=-1;

    for(int j=0; j<len; j++){
    	if(str[j]>='a' && str[j]<='z'){
    		swap(str[j],str[++i]);
    	}
    }
}

void proc2(string &str){
	int len=str.size();

	int p=0;
	int q=len-1;

	while(p<q){
		if(str[p]>='a' && str[p]<='z'){
			p++;
		}else if(str[q]>='A' && str[q]<='Z'){
			q--;
		}else{
			swap(str[p],str[q]);
		}
	}
}
int main(){

    string s="BabFaaCDE";

	//proc1(s);
	proc2(s);
	cout<<s<<endl;
}