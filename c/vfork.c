/*************************************************************************
	> File Name: vfork.c
	> Author: 
	> Mail: 
	> Created Time: Thu 12 Feb 2015 10:12:22 PM CST
 ************************************************************************/

#include <unistd.h>
#include <stdio.h>
 int main()
 {
          pid_t pid;
          int count = 0;
          pid = vfork();
          ++count;
          printf("count = %d\n",count);
          return 0;
      
 } 
