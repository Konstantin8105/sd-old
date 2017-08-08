package utils

import (
	"sort"
)

// UniqueInt - create unique int`s in array
func UniqueInt(array *[]int) {
	// sorting array of point index
	sort.Sort(sort.IntSlice(*array))
	// remove same point index
	/*
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
	*/
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

/*
func UniquePointIndex(points *[]point.Index) {
	ints := make([]int, len(*points), len(*points))
	for i := 0; i < len(*points); i++ {
		ints[i] = int((*points)[i])
	}
	UniqueInt(&ints)
	(*points) = (*points)[0:len(ints)]
	for i := 0; i < len(*points); i++ {
		(*points)[i] = point.Index(ints[i])
	}
}
*/
