package main

// =============================================================
// LogicalExpressions - GO
// main.go
// =============================================================
// Created by Lance T. Walker on 1/19/2019
// Copyright © 2019 CodeLife-Productions. All rights reserved.
// =============================================================

import (
	"fmt"
	// packages:
	Services "./services"
	// Utilities "./utils"
)

/*** System Properties ***/
const (
	VERSION string = "1.0.0"
)

// const Algorithms = Services.Algorithms;

/*** System Methods ***/
func fib(n int) int {
	fn := make(map[int]int)
	for i := 0; i <= n; i++ {
		var f int
		if i <= 2 {
			f = 1
		} else {
			f = fn[i-1] + fn[i-2]
		}
		fn[i] = f
		println("new fib: ", f)
	}
	return fn[n];
}

// ===> Entry:
func main() {
	fmt.Println("main.go: Starting")

	array := []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
	fmt.Println("initial array: ", array)
	newArrary := Services.Bubblesort(array)
	fmt.Println("final arrary: ", newArrary)

	strings := []string{"racecar", "one", "two", "mom", "1221", "121121"}
	for _, str := range strings {
		fmt.Println("RESULT OF TEST: ", str, isPalindrome(str))
	}

	// fib
	var x = fib(500)
	println(x)
}
