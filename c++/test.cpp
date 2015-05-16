/*************************************************************************
	> File Name: test.cpp
	> Author: 
	> Mail: 
	> Created Time: Wed 11 Feb 2015 10:50:35 PM CST
 ************************************************************************/

#include<iostream>
#include<cstring>
using namespace std;

class A {
private:
    int value;
public:
    A(int n){
        value=n;
        
    }
    A(A &other){
        value=other.value;
    }

    void Print(){
        cout<<value<<endl;
    }
};

int main(){
    A *a = new A(10);
    A *b=a;
    b->Print();

    A aa(100);
    aa.Print();
    //string test
    char str[10];

    strcpy(str,"012345678910111213");
    cout<<str<<endl;

    //string memory test

    char s1[] = "hello world";
    char s2[] = "hello world";

    char *s3 = "hello world";
    char *s4 = "hello world";

    cout<<(s1==s2)<<endl<<(s3==s4)<<endl;

    int n=9;
    string s(n,'0');

    cout<<s<<endl;
    for(int i=0; i<n; i++){
        s[i]+=i+1;
    }
    cout<<s<<endl;

    s.erase(s.begin());

    cout<<s<<endl;


}

