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
	initial, final []int
}

type Sort interface { // Sort Types
	bubble_sort() []int
	heap_sort() []int
	insertion_sort() []int
	merge_sort() []int
	quick_sort() []int
	selection_sort() []int
	shell_sort() []int
	topological_sort() []int
}

/*** Methods ***/
func algorithm() {

}

/*** Definitions ***/
// func (ent Result) bubble_sort() []int {

// }

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
