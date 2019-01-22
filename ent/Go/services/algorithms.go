package services

// =============================================================
// LogicalExpressions - GO
// algorithms.go
// =============================================================
// Created by Lance T. Walker on 1/19/2019
// Copyright © 2018 CodeLife-Productions. All rights reserved.
// =============================================================

// | Selection   | n^2                   | n^2              | O time complex                                        |
// | Insertion   | n^2                   | n^2              | product build one increment per operation             |
// | Bubble      | n^2                   | n^2              | swap pairs in list upon full incrementation of input  |
// | Quick       | n*log(n)              | n^2              |  O efficient, ordering list                           |
// | Heap        | n*log(n)              | n*log(n)         | sorts heap binary tree by comparing size of max input |
// | Merge       | n*log(n)              | n*log(n)         | divides complex input to smaller subset problems      |
// | Shell       | n*log(n)^2 or n^(3/2) | n/a              | complex application of general sort algorithm         |
// | Topological | n/a                   | n/a              | linear order of input vertices

/*** Properties ***/
type Types struct {
	Str []string // string arrary     {}=Nullable
	Int []int    // integer arrary    {}=Nullable
}

/*** Methods ***/
// func Selectionsort(ent Types) Types {

// }
// func Insertionsort(ent Types) Types {

// }
func Bubblesort(ent []int) []int {
	compl := true
	for compl {
		compl = false
		for i := 0; i < len(ent)-1; i++ {
			if ent[i+1] < ent[i] {
				tmp := ent
				ent[i], ent[i+1] = tmp[i+1], tmp[i]
				compl = true
			}
		}
	}
	return ent
}

// func Quicksort(ent Types) Types {

// }
// func Heapsort(ent Types) Types {

// }
// func Mergesort(ent Types) Types {

// }
// func Shellsort(ent Types) Types {

// }
// func Topologicalsort(ent Types) Types {

// }
