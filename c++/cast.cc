#include<iostream>
using namespace std;


int main() {
   float f =1.1f;
   int a = static_cast<int>(f);
   int b = -1;
   cout << static_cast<unsigned>(b);
   const char* cc = "hello";
   char* c = const_cast<char*>(cc);

}
