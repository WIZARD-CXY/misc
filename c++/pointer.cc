#include<iostream>
using namespace std;

void func(int *a){
   *a=10;
}
int main() {
   int a;
   cout<<a<<endl;
   func(&a);
   const int* rawpa = new int(11);
   shared_ptr<const int> pa = nullptr;
   pa = shared_ptr<const int>(rawpa);
   shared_ptr<const int> pb = pa;
   
   cout<<*pb<<" "<<pb.use_count()<<endl;
   pb.reset();
   cout<<" "<<pa.use_count()<<endl;
   cout<<"after "<<*rawpa; 
   pa.reset();
   cout<<"after "<<*rawpa; 
}
   
