#include<iostream>
#include<thread>
using namespace std;
void func() { std::cout << "new thread \n"; }
int main() {
    std::thread t(func);
    if (t.joinable()) {
        t.join(); // 或者t.detach(); 
    }
    return 0;
}
