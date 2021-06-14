#include<iostream>
using namespace std;
namespace n1 {
void func() {
  std:cout<<"n1 func"<<endl;
}
}

namespace n2 {
void func() {
   std:cout<<"n2 func"<<endl;
}
}

int main() {
   n1::func();
   n2::func();
}
