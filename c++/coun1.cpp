#include <iostream>
#include <cmath>
#include <vector>
using namespace std;

int main()
{
    int n = 1021;
    int num = 0;
    int tmp = n;
    vector<int> pos;
    while (n)
    {
        n &= (n - 1);

        num++;
        pos.push_back(log2(tmp - n) + 1);
        tmp = n;
    }

    cout << num << endl;
    for (int i = 0; i < pos.size(); i++)
    {
        cout << pos[i] << " ";
    }
}