#include <memory>
#include <iostream>
using namespace std;

void fun(int *a) {
   *a=9;
}
int main() {
  shared_ptr<int> a = make_shared<int>(19);
  cout<<*a<<endl;
  fun(a.get());
  cout<<*a<<endl;
}

