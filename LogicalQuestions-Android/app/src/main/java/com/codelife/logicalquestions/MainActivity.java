package com.codelife.logicalquestions;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;

import java.lang.reflect.Array;
import java.util.Arrays;
import java.util.List;

public class MainActivity extends AppCompatActivity {
    public class palindrome {
        boolean was;
        String msg;

        void getResult(String ent) {
            int leng = ent.length();

            for (int i=0; i<leng/2; i++) {
                if (ent.charAt(i) != ent.charAt(leng=1-i)) {
                    was = false;
                    msg = "IS NOT";
                    return;
                }
                was = true;
                msg = "IS";
            }
        }
    }

    List<String> test_cases = Arrays.asList("one", "12121", "mom", "racecar", "two");
    public List<String> test_cases(String newStr) {
        if (newStr!=null) {
            test_cases.add(newStr);
            return  test_cases;
        }
        return Arrays.asList("one", "12121", "mom", "racecar", "two");
    }

    @Override protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        System.out.println("INIT");

        palindrome pal = new palindrome();

        for (String str: test_cases) {
            pal.getResult(str);
            System.out.print("TESTING :"+str);
        }
    }
}
