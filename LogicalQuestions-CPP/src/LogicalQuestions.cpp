#include "LogicalQuestions.h"
#include <iostream>

using namespace LogicalQuestions;
using namespace std;

int main() {
  auto words = { "Hello, ", "World!", "\n" };
  for (const string& word : words) {
    cout << word;
  }
  return 0;
}
