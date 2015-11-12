void do_it();
int main()
   {
           do_it();
           return 0;
   }
   void do_it()
   {
           char* p = 1; //定义一个字符指针变量a，指向地址1，这个地址肯定不是自己可以访问的，但是这行不会产生段错误
          *p = 'a'; //真正产生段错误的在这里，试图更改地址1的值，此时内核会终止该进程，并且把core文件dump出来
  }

 $ ulimit -c unlimited #enable with this command