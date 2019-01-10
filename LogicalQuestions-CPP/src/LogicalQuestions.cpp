#include <algorithm>
#include <array>
#include <iostream>
#include <string>
#include <vector>

using namespace std;

string test_cases[] = {"one", "12121", "mom", "racecar", "two"};

struct palindrome {
  bool was;
  string msg;
};

static palindrome getResult(string ent) {
  palindrome result;
  result.was = true;
  result.msg = "IS";
  size_t leng = ent.length();

  for (int i = 0; i < leng / 2; i++) {
    if (ent[i] != ent[leng - 1 - i]) {
      result.was = false;
      result.msg = "IS NOT";
    }
  }
  return result;
}
///////////////////
static int binarySearch(int arr[], int l, int r, int x) {
  int mid = l + (r - l) / 2;

  if (r >= l) {
    if (arr[mid] == x)
      return mid;
    if (arr[mid] > x)
      return binarySearch(arr, l, mid - 1, x);
    return binarySearch(arr, mid + 1, r, x);
  }
  return -1;
}

int main(int argc, const char *arg[]) {
  cout << "INIT\n";

  for (int i = 0; i < 5; i++) {
    string str = test_cases[i];
    palindrome res = getResult(str);

    cout << "TESTING: " << str << " " << res.msg << " A PALINDROME.\n";
  }

  int arr[] = {2, 3, 4, 10, 40};
  int x = 10;
  int n = sizeof(arr) / sizeof(arr[0]);

  int result = binarySearch(arr, 0, n - 1, x);
  (result == -1) ? cout << "QUERY IS NOT PRESENT IN ARRARY"
                 : cout << "QUERY IS PRESENT AT INDEX: " << result;

  return 0;
}

// void insertion_sort(item s[], int n) {
//   int i,j; // counters

//   for (i=1; i<n; i++) {
//     j=i;
//     while ((j>0)&&(s[j] < s[j-1]]) {
//       swap (&s[j], &s[j-1]);
//       j=j-1;
//     }
//  }
// }
