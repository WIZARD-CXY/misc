#include<iostream>
using namespace std;

class B {
public:
      B();
      B(int i);
      ~B();
      void print() const;
private:
      int b;
};


B::B() {
     b=0;
     cout<<"B default constructor is called"<<endl;
}


B::B(int i) {
    b = i;
    cout<<"B's int constructor called"<<endl;
}

B::~B(){
    cout<<"B's destructor called"<<endl;
}

void B::print() const {
    cout<<b<<endl;
}

class C : public B {
public:
    C();
    C(int i, int j);
    ~C();
    void print() const;
private:
    int c;
};

C::C(){
   cout<<"C constructor called"<<endl;
}
C::C(int i, int j): B(i), c(j){
   cout<<"c int constructor called"<<endl;
}

C::~C() {
   cout<<"c destructor called"<<endl;
}

void C::print() const {
     B::print();
     cout<<c<<endl;
}

int main(){
    C obj(5,6);
    obj.print();
}
