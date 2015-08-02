/*************************************************************************
	> File Name: virtual.cpp
	> Author: 
	> Mail: 
	> Created Time: Fri 13 Feb 2015 09:39:51 PM CST
 ************************************************************************/

#include<iostream>
using namespace std;

class Person{
    int id;
    char name[100];
public:
    virtual void aboutMe(){
        cout<<"I'm a person"<<endl;
    }
    virtual bool addCourse(string s)=0;
    virtual ~Person(){
        cout<<"deleting a person"<<endl;
    }
};

class Student :public Person{
public:
    void aboutMe(){
        cout<<"I am a student"<<endl;
    }

    bool addCourse(string s){
        cout<<"Add "<<s<<" to student"<<endl;
        return true;
    }
    ~Student(){
        cout<<"deleting a student"<<endl;
    }

};

int main(){

    Person *p =new Student();
    p->aboutMe();
    p->addCourse("math");

    delete p;

    int a=12;
    int &b=a;
    cout<<b<<endl;
    int c=111;
    b=c;
    cout<<b<<endl;
}


