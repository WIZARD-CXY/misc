/*************************************************************************
	> File Name: functionptr.c
	> Author: 
	> Mail: 
	> Created Time: Sat 14 Feb 2015 07:04:25 PM CST
 ************************************************************************/

#include<stdio.h>
#include<assert.h>

typedef int (*FP_CALC)(int,int); //define a function type pointer Named FP_CALC

int add(int a, int b){
    return a+b;
}

int sub(int a, int b){
    return a-b;
}

int mul(int a, int b){
    return a*b;
}

int div(int a, int b){
    return b?a/b:-1;
}

//define a function return a pointer which point to the coresponding function.
FP_CALC calc_func(char op){
    switch(op){
        case '+':
        return add;
        case '-':
        return sub;
        case '*':
        return mul;
        case '/':
        return div;
        default:
        return NULL;
    }
}

//define a function pointer again 
//
int (*s_calc_func(char op)) (int ,int){
    return calc_func(op);
}

int calc(int a, int b, char op){
    FP_CALC fp = calc_func(op);
    int (*s_fp)(int,int) = s_calc_func(op);

    assert(fp==s_fp);

    if(fp){
        return fp(a, b);
    }else{
        return -1;
    }
}

void main(){
    int a=100, b=20;

    printf("calc(%d, %d, %c) = %d\n", a, b, '+', calc(a, b, '+'));
    printf("calc(%d, %d, %c) = %d\n", a, b, '-', calc(a, b, '-'));
    printf("calc(%d, %d, %c) = %d\n", a, b, '*', calc(a, b, '*'));
    printf("calc(%d, %d, %c) = %d\n", a, b, '/', calc(a, b, '/'));
}
