package main

// =============================================================
// LogicalExpressions - GO
// main.go
// =============================================================
// Created by Lance T. Walker on 1/19/2019
// Copyright © 2018 CodeLife-Productions. All rights reserved.
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
// ===> Entry:
func main() {
	fmt.Println("main.go: Starting")

	array := []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
	fmt.Println("initial array: ", array)
	newArrary := Services.Bubblesort(array)
	fmt.Println("final arrary: ", newArrary)

	// compl := true;
	// fmt.Println("initial arrary: ", array)
	// for compl {
	//   compl = false;
	//   for i := 0; i < len(array) - 1; i++ {
	//     if array[i + 1] < array[i] {
	//       tmp := array;
	//       array[i], array[i + 1] = tmp[i + 1], tmp[i];
	//       compl = true;
	//     }
	//   }
	// }
	// fmt.Println("final array: ", array)
}
