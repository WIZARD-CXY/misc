#include<iostream>
#include<memory>
using namespace std;

class TClass {
private:
      int a_;
      int *data_;
public:
      TClass(){
        data_ = new int[10];
      }

      TClass(int a): a_(a) {
        cout<<"default constructor2"<<endl;
        data_ = new int[10];
      }
      
      TClass(const TClass &a): a_(a.a_) {
        cout<<"copy constructor"<<endl;
        data_ = new int[10];

        memcpy(data_, a.data_, 10*sizeof(int));
      }

      TClass(TClass &&a): a_(a.a_), data_(a.data_){
         a.data_ = nullptr;
         cout<<"move constructor"<<"\n";
      }

      TClass& operator=(TClass &a) {
          a_ = a.a_;
          cout<<"data_ "<< data_<<endl;
          if (data_) delete[] data_;
          data_ = new int[10];
          memcpy(data_, a.data_, 10*sizeof(int));

          cout<<"assign constructor"<<"\n";
          return *this;
      }

      TClass& operator=(TClass &&a) {
          a_ = a.a_;
          if(data_) delete[] data_;
          data_ = a.data_;
          a.data_ = nullptr;
          cout<<"move assign func"<<"\n";
          return *this;
      }
    
      ~TClass(){
          cout<<"deconstructor"<<"\n";
          if(data_) delete[] data_;
      }

      void func() {
         cout<< "a " <<a_<<"\n";
      }
      
};

int main(){
   TClass a;
   TClass a1(1);
   TClass b(a1);
   TClass c(std::move(a1));
   c = b;
   c = std::move(b);
   TClass d = c;
   TClass e = std::move(d);
   a.func();
   b.func();

   return 0;
}
