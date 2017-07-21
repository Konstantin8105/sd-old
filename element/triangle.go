package element

import (
	"fmt"

	"github.com/Konstantin8105/GoFea/point"
)

// Triangle - property of triangle element
type Triangle struct {
	Index        ElementIndex
	PointIndexes [3]point.Index
}

func (e Triangle) ElementDescription() string {
	return fmt.Sprintf("Triangle element №%v with points %v", e.Index, e.PointIndexes)
}
