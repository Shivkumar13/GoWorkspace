#include <iostream>
#include <vector>
using namespace std;

int main() {

	    vector<int> arr = {4, 7, 8, 3, 8, 67, 34, 23, 555, 6563, 3465, 4, 32, 23, 456, 43};
	    int n = arr.size();
        int i, key;
        int maxSoFar = -1000000;
        int minSoFar = 2000000;

        for (i =0;i < n; i++){
            maxSoFar = max(maxSoFar, arr[i]);
            minSoFar = min(minSoFar, arr[i]);
        }


        cout << "Max element: "<< maxSoFar << endl;
        cout << "Min element: "<< minSoFar << endl;

         return 0;
}

