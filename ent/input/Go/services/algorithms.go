package services
// =============================================================
// LogicalExpressions - GO
// algorithms.go
// =============================================================
// Created by Lance T. Walker on 1/19/2019
// Copyright © 2019 CodeLife-Productions. All rights reserved.
// =============================================================

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
