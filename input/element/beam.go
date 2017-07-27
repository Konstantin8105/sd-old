package element

import (
	"github.com/Konstantin8105/GoFea/input/point"
)

// Beam - property of beam element
type Beam struct {
	index        Index
	pointIndexes []point.Index
}

// GetIndex - return index of beam
func (b Beam) GetIndex() Index {
	return b.index
}

// GetPointIndex - return indexes of point for that finite element
func (b Beam) GetPointIndex() []point.Index {
	return b.pointIndexes
}

// GetAmountPoint - return amount points in finite element
func (b Beam) GetAmountPoint() int {
	return 2
}
