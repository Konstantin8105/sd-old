package utils

import "sort"

func UniqueInt(array *[]int) {
	// sorting array of point index
	sort.Sort(sort.IntSlice(*array))
	// remove same point index
	amount := 0
	for i := range *array {
		if i == 0 {
			amount++
			continue
		}
		if (*array)[i-1] != (*array)[i] {
			amount++
		}
	}
	inx := 0
	for i := range *array {
		if i == 0 {
			inx++
			continue
		}
		if (*array)[i-1] != (*array)[i] {
			(*array)[inx] = (*array)[i]
			inx++
		}
	}
	(*array) = (*array)[0:inx]
}
