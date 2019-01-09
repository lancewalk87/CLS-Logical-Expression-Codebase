using System;

namespace LogicalQuestions {
    class Program {
        private static string isPalindrome(string ent)
        { // String Check
            int leng = ent.Length; 

            for (int i=0; i<leng/2; i++)
            {
                if (ent[i] != ent[leng - 1 - i]) return "IS NOT"; 
            }
            return "IS"; 
        }

        static void Main(string[] args) {
            Console.WriteLine("INIT: PALINDROME");
            string[] test_cases = new string[] { "one", "two", "12121", "racecar", "mom", "test" };

            // index: test_cases 
            foreach (var i in test_cases) Console.WriteLine("TESTING: {1}, {1} A PALINDROME.", i, isPalindrome(i));
        }
    }
}
