package dof

import (
	"sort"

	"github.com/Konstantin8105/GoFea/element"
	"github.com/Konstantin8105/GoFea/point"
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
	dofArray  []int
	dimension Dim
}

// NewBeam - add new beam
func NewBeam(beams []element.Beam, dim Dim) (d DoF) {
	array := make([]int, len(beams)*2, len(beams)*2)
	for i := range beams {
		array[i*2+0] = int(beams[i].PointIndexes[0])
		array[i*2+1] = int(beams[i].PointIndexes[1])
	}
	utils.UniqueInt(&array)
	d.dofArray = array
	d.dimension = dim
	return d
}

// GetDoF - get degree of freedom for point index
func (d *DoF) GetDoF(index point.Index) []AxeNumber {
	if d.dimension == Dim2d {
		axes := make([]AxeNumber, int(d.dimension), int(d.dimension))
		number := d.dofArray[d.found(index)]
		for i := 0; i < int(d.dimension); i++ {
			axes[i] = AxeNumber(number + i*len(d.dofArray))
		}
		return axes
	}
	empty := make([]AxeNumber, 0, 0)
	return empty
}

func (d *DoF) found(index point.Index) int {
	return sort.SearchInts(d.dofArray, int(index))
}
