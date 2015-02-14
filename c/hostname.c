#include <stdio>

int child_main(void* arg)
{
  printf(" - World !\n");
  sethostname("In Namespace", 12);
  execv(child_args[0], child_args);
  printf("Ooops\n");
  return 1;
}
 
int main()
{
  printf(" - Hello ?\n");
  int child_pid = clone(child_main, child_stack+STACK_SIZE,
      CLONE_NEWUTS | SIGCHLD, NULL);
  waitpid(child_pid, NULL, 0);
  return 0;
}
