#include<iostream>
using namespace std;
template<typename T>
T mymax(T a, T b) {
   return a > b ? a:b;
}

template <typename T>
class Vec {
public:
    Vec(T a): a_(a){}

    void func() { cout <<"f value:"<<a_<<endl;}

    T a_;
};

int main() {
   cout<<mymax(1,2)<<endl;
   cout<<mymax(1.1f,2.2f)<<endl;


   Vec<int> vi(1);
   vi.func();

   Vec<float> vf(1.1f);
   vf.func();
}
