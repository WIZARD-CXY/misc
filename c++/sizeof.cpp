/*************************************************************************
	> File Name: sizeof.cpp
	> Author: 
	> Mail: 
	> Created Time: Tue 24 Feb 2015 11:45:01 AM CST
 ************************************************************************/

#include<cstdio>
#include<iostream>
#include<string>
using namespace std;

class A{
    public:
    A(){}
    virtual ~A(){}
};


struct s1
{
char a[8];
};

struct s2
{
double d;
};

struct s3
{
s1 s;
char a;
};

struct s4
{
s2 s;
char a; 
};
int main(){
    printf("sizeof empty {} is %d\n",sizeof(A));

    string s="helllo world     hahahhahahah";

    printf("%d\n",sizeof(s));

    cout<<sizeof(s1)<<endl; // 24
	cout<<sizeof(s2)<<endl; // 16
	cout<<sizeof(s3)<<endl; // 16
	cout<<sizeof(s4)<<endl; // 16
}
