/*************************************************************************
	> File Name: sizeof.c
	> Author: 
	> Mail: 
	> Created Time: Tue 24 Feb 2015 11:45:01 AM CST
 ************************************************************************/

#include<cstdio>
class A{
    public:
    A(){}
    virtual ~A(){}
};

int main(){
    printf("sizeof empty {} is %d\n",sizeof(A));
}
