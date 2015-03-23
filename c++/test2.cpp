#include <stdio.h>

int main(void)
{
	double i;
	int j;
	
	for (i = 0, j = 0; i != 10.0 && j != 102; i += 0.1, j++) {
		printf("%.50f %d\n", i,j);
	}		
   	return 0;
}
