using System;

namespace LogicalQuestions {
    class Program {
        struct palindrome
        {
            public bool was;
            public string msg; 
        }

        static palindrome getResult(string ent)
        {
            palindrome result = new palindrome { was = true, msg = "IS" };
            int leng = ent.Length;
            
            for (int i=0; i<leng/2; i++)
            {
                if (ent[i] != ent[leng - 1 - i]) { result.was = false; result.msg = "IS NOT"; } 
            }

            return result; 
        }

       //  palindrome* test_results = null; 

        static void Main(string[] args) {
            Console.WriteLine("INIT: PALINDROME");
            
            string[] test_cases = new string[] { "one", "12121", "mom", "racecar", "two" };

            // index: test_cases:
           foreach (var i in test_cases)
            {
                palindrome res = getResult(i);
                Console.WriteLine("TESTING: {0}, {1} A PALINDROME", i, res.msg); 
            }
        }
    }
}
