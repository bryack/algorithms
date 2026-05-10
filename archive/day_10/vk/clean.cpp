#include <iostream>
#include <vector>

using namespace std;

int main() {
    int size;
    cin >> size;
    
    if (size == 0) {
        return 0;
    }
    
    vector<int> arr(size + 1);
    for (int i = 0; i < size + 1; i++) {
        cin >> arr[i];
    }
    
    int element = arr[size];
    bool first = true;
    
    for (int i = 0; i < size + 1; i++) {
        if (arr[i] != element) {
            if (!first) {
                cout << " ";
            }
            cout << arr[i];
            first = false;
        }
    }
    
    return 0;
}
