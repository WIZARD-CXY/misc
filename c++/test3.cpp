#include <stdio.h>

int main(){
	int a,b,c;

	scanf("%d%d%d",&a,&b,&c);

	printf("%.*f\n", c, double(a)/b);
}