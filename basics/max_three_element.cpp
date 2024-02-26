#include <iostream>
#include <bits/stdc++.h>
#include <vector>
using namespace std;

int main()
{
    vector<int> arr = {8,34,2,22,15};

    int i = 0, j = 0, key, n = arr.size();

    int maxSoFar = -100000000, maxSoFar2= -200000, maxSoFar3=-3000000;

    for(i = 0; i < n; i++){
        cout << "Given Array: " << arr[i] << endl;
    }

    for(i = 0; i < n; i++){
        key = arr[i];

        for(j = i+1; j < n; j++){
            if(arr[j] > key){
                arr[j+1] = arr[j];
            }
            arr[j] = key;
        }
    }

   for(i = 0; i < n; i++){
        cout << "Given Array: " << arr[i] << endl;
    }
}