package dof

import (
	"fmt"
	"sort"

	"github.com/Konstantin8105/GoFea/input/point"
	"github.com/Konstantin8105/GoFea/utils"
)

// Dim - dimension unit
type Dim int

// Type of dimension
const (
	Dim2d Dim = 3 // 3 degree of freedom for point in 2d. Dx, Dy, M
	Dim3d     = 6 // 6 degree of freedom for point in 3d. Dx, Dy, Dz, Mx, My, Mz
)

// AxeNumber - axe of number
type AxeNumber int

// DoF - degree of freedom
type DoF struct {
	DofArray  []point.Index
	Dimension Dim
}

// GetDoF - get degree of freedom for point index
func (d *DoF) GetDoF(index point.Index) []AxeNumber {
	axes := make([]AxeNumber, int(d.Dimension), int(d.Dimension))
	number := d.found(index)
	for i := 0; i < int(d.Dimension); i++ {
		axes[i] = AxeNumber(i + number*int(d.Dimension))
	}
	return axes
}

func (d *DoF) found(index point.Index) int {
	i := sort.Search(len(d.DofArray), func(i int) bool { return int(d.DofArray[i]) >= int(index) })
	if i >= 0 && i < len(d.DofArray) && int(d.DofArray[i]) == int(index) {
		// index is present at array[i]
		return i
	}
	// index is not present in array,
	// but i is the index where it would be inserted.
	panic("Not correct binary searching")
}

// RemoveIndexes - remove indexex for axeNumber slice
// without reallocation matrix
func RemoveIndexes(a *[]AxeNumber, indexes ...int) {
	if len(indexes) == 0 {
		return
	}
	// sorting indexes for optimization of algoritm
	utils.UniqueInt(&indexes)
	UniqueAxeNumber(a)

	// global checking indexes
	if indexes[len(indexes)-1] >= len(*a) || indexes[0] < 0 {
		panic(fmt.Errorf("indexes is outside of matrix. Indexes = %v", indexes))
	}
	// modify values
	positionIndex := 0
	newPositionInSlice := 0
	for i := 0; i < len(*a); i++ {
		if positionIndex != len(indexes) && i == indexes[positionIndex] {
			positionIndex++
			continue
		}
		(*a)[newPositionInSlice] = (*a)[i]
		newPositionInSlice++
	}

	(*a) = (*a)[0 : len(*a)-len(indexes)]
}

// UniqueAxeNumber - create unique axes in array
func UniqueAxeNumber(axes *[]AxeNumber) {
	ints := make([]int, len(*axes), len(*axes))
	for i := 0; i < len(*axes); i++ {
		ints[i] = int((*axes)[i])
	}
	utils.UniqueInt(&ints)
	(*axes) = (*axes)[0:len(ints)]
	for i := 0; i < len(*axes); i++ {
		(*axes)[i] = AxeNumber(ints[i])
	}
}
