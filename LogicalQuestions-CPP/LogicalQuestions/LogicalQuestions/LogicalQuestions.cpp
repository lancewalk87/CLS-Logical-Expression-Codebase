#include "pch.h"
#include <iostream>
#include <string>

using namespace std;

string test_cases[] = { "one", "12121", "mom", "racecar", "two" };
struct palindrome {
	bool was;
	string msg;
};

palindrome getResult(string ent) {
	palindrome result;  result.was = true, result.msg = "IS";
	size_t leng = ent.length();

	for (int i = 0; i < leng / 2; i++) {
		if (ent[i] != ent[leng - 1 - i]) result.was = false, result.msg = "IS NOT";
	}
	return result;
}

palindrome *test_results = NULL;

int main() {
	cout << "INIT\n";
	test_results = new palindrome[ARRAY_SIZE(test_cases)];

	// one:
	for (int i = 0; i < 5; i++) {
		string str = test_cases[i];
		palindrome res = getResult(str);

		cout << "TESTING: " << str << " " << res.msg << " A PALINDROME.\n";
		test_results[i] = res;
	}
	cout << "COMPLETE\n with exit size: \n " << ARRAY_SIZE(test_results);
}
