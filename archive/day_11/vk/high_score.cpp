#include <iostream>
#include <vector>
using namespace std;

int main() {
    int n;
    cin >> n;
    
    if (n == 0) {
        return 0;
    }
    
    vector<int> arr(n);
    for (int i = 0; i < n; i++) {
        cin >> arr[i];
    }
    
    int maxScore = arr[0];
    for (int i = 1; i < n; i++) {
        if (arr[i] > maxScore) {
            maxScore = arr[i];
        }
    }
    
    cout << maxScore;
    return 0;
}
