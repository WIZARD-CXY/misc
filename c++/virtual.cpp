/*************************************************************************
	> File Name: virtual.cpp
	> Author: 
	> Mail: 
	> Created Time: Fri 13 Feb 2015 09:39:51 PM CST
 ************************************************************************/

#include<iostream>
using namespace std;

class Base1{
public:
    virtual void f(){cout<<"Base1::f"<<endl;}
    virtual void g(){cout<<"Base2::g"<<endl;}
    virtual void h(){cout<<"Base3::h"<<endl;}

};

class Base2{
    virtual void f(){cout<<"Base2::f"<<endl;}
    virtual void g(){cout<<"Base2::g"<<endl;}
    virtual void h(){cout<<"Base2::h"<<endl;}
};

class Base3{
    virtual void f(){cout<<"Base3::f"<<endl;}
    virtual void g(){cout<<"Base3::g"<<endl;}
    virtual void h(){cout<<"Base3::h"<<endl;}
};

class Derive : public Base1, public Base2, public Base3{
public:
    virtual void f(){cout<<"Derive::f"<<endl;}
    virtual void g1(){cout<<"Derive::g1"<<endl;}

};

typedef void(*Fun)(void);

int main(){
    Fun pFun = NULL;

    Derive d;
    int** pVtab = (int**)&d;

    //Base1's vtable
    pFun = (Fun)pVtab[0][0];
    pFun();

//    pFun = (Fun)pVtab[0][1];
    pFun = (Fun)*((int*)*(int*)((int*)&d+0)+1);
    pFun();

    pFun = (Fun)pVtab[0][2];
    pFun();
    
    pFun = (Fun)pVtab[0][3];
    pFun();

    
//    pFun = (Fun)pVtab[0][4];
//    pFun();

    //Base2's vtable
    pFun = (Fun)pVtab[1][0];
    pFun();

    pFun = (Fun)pVtab[1][1];
    pFun();

    pFun = (Fun)pVtab[1][2];
    pFun();
    
    pFun = (Fun)pVtab[1][3];
    pFun();

    //Base3's vtable
    pFun = (Fun)pVtab[2][0];
    pFun();

    pFun = (Fun)pVtab[2][1];
    pFun();

    pFun = (Fun)pVtab[2][2];
    pFun();
    
    pFun = (Fun)pVtab[2][3];
    pFun();

    return 0;
}


