#include <iostream>
using namespace std;

int abs(int x){
	int sign=x>>31; //sign =0 or -1
	return (x^sign)-sign;
}
int main(){
	int x;
	cin>>x;

	cout<<abs(x)<<endl;
}