package finiteElement

import (
	"github.com/Konstantin8105/GoFea/dof"
	"github.com/Konstantin8105/GoLinAlg/matrix"
)

// RemoveZeros - remove columns, rows of matrix and columns of dof
func RemoveZeros(matrix *matrix.T64, axes *[]dof.AxeNumber) {
	var removePosition []int
	// TODO: len --> to matrix lenght
	// TODO: at the first check diagonal element
	for i := 0; i < len(*axes); i++ {
		found := false
		for j := 0; j < len(*axes); j++ {
			if (*matrix).Get(i, j) != 0.0 {
				found = true
				break
			}
		}
		if found {
			continue
		}
		removePosition = append(removePosition, i)
	}

	// TODO: can parallel
	// remove row and column from global stiffiner
	(*matrix).RemoveRowAndColumn(removePosition...)
	// remove column from axes
	dof.RemoveIndexes(axes, removePosition...)
}
