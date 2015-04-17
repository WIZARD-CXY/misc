#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/ipc.h>
#include <fcntl.h>
#include <sys/stat.h>

int main(){
	int fd1;

	if((fd1=open("./file1",O_RDWR))==-1){
		perror("open file failed haha");
	}

	pid_t pid=fork();

	if(pid==0){
		sleep(1);
		int fd2;
		if((fd2=open("./file1",O_RDWR))==-1){
			perror("open failed file2 haha");
		}

		printf("fd1=%d\tfd2=%d\n",fd1,fd2);
		lseek(fd1,0,SEEK_SET);
		if(write(fd1,"a",1)==-1){
			perror("write failed fd1");
		}

		lseek(fd2,1,SEEK_SET);
		if(write(fd2,"b",1)==-1){
			perror("write failded fd2");
		}

		close(fd1);
		close(fd2);
	} else {
		printf("fd1=%d\n",fd1);
		lseek(fd1,0,SEEK_SET);
		if(write(fd1,"c",1)==-1){
			perror("write fd1 again");
		}
		close(fd1);
	}
}