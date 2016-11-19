#include <cstdio>
using namespace std;
const int mod = 1e6;
int main()
{
    int sum = 0;
    int factorial = 1;
    for (int i = 1; i <= 40; i++)
    {
        factorial *= i;
        sum += factorial;
    }

    printf("%d\n", sum);

    sum = 0;
    for (int i = 1; i <= 40; i++)
    {
        factorial = 1;
        for (int j = 1; j <= i; j++)
        {
            factorial = (factorial * j) % mod;
        }
        sum = (sum + factorial) % mod;
    }
    printf("%d\n", sum);
}