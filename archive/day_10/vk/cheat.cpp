#include <iostream>
#include <vector>

using namespace std;

int main() {
    int size;
    cin >> size;
    
    if (size == 0) {
        return 0;
    }
    
    vector<int> arr(size);
    for (int i = 0; i < size; i++) {
        cin >> arr[i];
    }
    
    int d = 0;
    for (int i = 0; i < size; i++) {
        if (arr[i] != 0) {
            arr[d] = arr[i];
            d++;
        }
    }
    
    for (int i = d; i < size; i++) {
        arr[i] = 0;
    }
    
    for (int i = 0; i < size; i++) {
        if (i == size - 1) {
            cout << arr[i];
        } else {
            cout << arr[i] << " ";
        }
    }
    
    return 0;
}
