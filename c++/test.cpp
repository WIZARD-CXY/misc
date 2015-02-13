/*************************************************************************
	> File Name: test.cpp
	> Author: 
	> Mail: 
	> Created Time: Wed 11 Feb 2015 10:50:35 PM CST
 ************************************************************************/

#include<iostream>
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
    A a = A(10);
    A b=a;
    b.Print();
}
