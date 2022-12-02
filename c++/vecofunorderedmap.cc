#include <vector>
#include <unordered_map>
#include <iostream>
using namespace std;

// Function to print vector elements
void print(vector<unordered_map<int, int> >& vect)
{

	cout << "vector : \n";
	for (int i = 0; i < (int)vect.size(); i++) {

		// Each element of the vector is a unordered map
		unordered_map<int, int> unorderedMap = vect[i];

		cout << "unordered map : ";
		cout << "[ ";

		// Print unordered map elements
		for (auto it = unorderedMap.begin();
			it != unorderedMap.end(); it++) {
			cout << it->first << ':' << it->second << " ";
		}

		cout << "]\n";
	}
}

// Driver Code
int main()
{

	// Declaring a vector of unordered maps
	vector<unordered_map<int, int> > vect;

	// Declaring a unordered map
	unordered_map<int, int> unorderedMap1;

	// Hashing values
	unorderedMap1[2] = 1;
	unorderedMap1[4] = 7;
	unorderedMap1[6] = 10;
        unorderedMap1[100] = 100;

	// Push back the unordered map in the vector
	vect.push_back(unorderedMap1);

	// Declaring another unordered map
	unordered_map<int, int> unorderedMap2;

	// Hashing values
	unorderedMap2[14] = 11;
	unorderedMap2[15] = 21;
	unorderedMap2[6] = 34;

	// Push back the unordered map in the vector
	vect.push_back(unorderedMap2);

	// Declaring another unordered map
	unordered_map<int, int> unorderedMap3;

	// Hashing values
	unorderedMap3[7] = 277;
	unorderedMap3[18] = 188;
	unorderedMap3[9] = 399;
	// Push back the unordered map in the vector
	vect.push_back(unorderedMap3);

	// Declaring another unordered map
	unordered_map<int, int> unorderedMap4;

	// Hashing values
	unorderedMap4[121] = 88;
	unorderedMap4[97] = 99;
	unorderedMap4[197] = 199;

	// Push back the unordered map in the vector
	vect.push_back(unorderedMap4);
        vect.resize(5);
	print(vect);
        cout<<"yahaha " << vect[0].size() << " " << vect[0].max_size() << endl;
	return 0;
}

