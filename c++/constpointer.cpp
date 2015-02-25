/*************************************************************************
	> File Name: constpointer.cpp
	> Author: 
	> Mail: 
	> Created Time: Tue 24 Feb 2015 11:53:59 AM CST
 ************************************************************************/
#include <iostream>

using namespace std;

int main(int argc, char *argv[])
{
    int a=3;
    int b;

    /*定义指向const的指针（指针指向的内容不能被修改）*/ 
    // const is left of *
    const int* p1; 
    int const* p2; 

    /*定义const指针(由于指针本身的值不能改变所以必须得初始化）*/ 
    //const is right of *
    int* const p3=&a; 

    /*指针本身和它指向的内容都是不能被改变的所以也得初始化*/
    const int* const p4=&a;
    int const* const p5=&b; 

    p1=p2=&a; //正确
    *p1=*p2=8; //不正确（指针指向的内容不能被修改）

    *p3=5; //正确
    p3=p1; //不正确（指针本身的值不能改变） 

    p4=p5;//不正确 （指针本身和它指向的内容都是不能被改变） 
    *p4=*p5=4; //不正确（指针本身和它指向的内容都是不能被改变） 

    return 0;
}
