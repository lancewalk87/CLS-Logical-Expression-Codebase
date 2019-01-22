package services

// =============================================================
// LogicalExpressions - GO
// algorithms.go
// =============================================================
// Created by Lance T. Walker on 1/19/2019
// Copyright © 2018 CodeLife-Productions. All rights reserved.
// =============================================================

/*** Properties: Algorithms ***/
type Result struct { // Defines entry, and final array
  // Sorted_MinMax
}

type Sort interface { // Sort Types
	bubble_sort() []int      // swap upon condition of invariant defined (i.e. range of numbers)
	heap_sort() []int        // sort binary tree based upon spec. comparison (i.e. range of numbers)
	insertion_sort() []int   //
	merge_sort() []int
	quick_sort() []int
	selection_sort() []int
	shell_sort() []int
	topological_sort() []int
}

/*** Methods ***/

/*** Definitions ***/
func Bubblesort(ent []int) []int {
  compl := true
  arr := ent;

  for compl {
    compl = false;
    for i := 1; i < len(arr) - 1; i++ {
      if arr[i + 1] < arr[i] {
        tmp := arr;
        arr[i], arr[i + 1] = tmp[i + 1], tmp[i]
        compl = true
      }
    }
  }
  return arr;
}

// func (ent Result) heap_sort() []int {

// }

// func (ent Result) insertion_sort() []int {

// }

// func (ent Result) merge_sort() []int {

// }

// func (ent Result) quick_sort() []int {

// }

// func (ent Result) selection_sort() []int {

// }

// func (ent Result) shell_sort() []int {

// }

// func (ent Result) topological_sort() []int {

// }
