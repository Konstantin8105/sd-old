package element

import (
	"fmt"

	"github.com/Konstantin8105/GoFea/point"
)

// Beam - property of beam element
type Beam struct {
	Index        ElementIndex
	PointIndexes [2]point.Index
}

func (e Beam) ElementDescription() string {
	return fmt.Sprintf("Beam element №%v with points %v", e.Index, e.PointIndexes)
}
