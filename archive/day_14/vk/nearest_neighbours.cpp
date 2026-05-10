#include <iostream>
#include <climits>

using namespace std;

int main() {
    int n;
    cin >> n;
    
    int prev;
    cin >> prev;
    
    int minDiff = INT_MAX;
    int bestX = 0, bestY = 0;
    
    for (int i = 1; i < n; i++) {
        int current;
        cin >> current;
        
        int diff = current - prev;
        if (diff < minDiff) {
            minDiff = diff;
            bestX = prev;
            bestY = current;
        }
        
        prev = current;
    }
    
    cout << bestX << " " << bestY << endl;
    
    return 0;
}
