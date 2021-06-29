#include<iostream>
using namespace std;

int main() {
   int x=10;
   

   auto add_x = [x](int a) { return x+a;};
   auto multiply_x = [&x](int a) { return a*x; };
   
   cout<<add_x(10)<<" "<<multiply_x(10)<<endl;
}
