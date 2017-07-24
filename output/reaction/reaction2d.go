package reaction

import "github.com/Konstantin8105/GoFea/input/point"

// Dim2 - reactions in 2d point with support
// Unit - N
type Dim2 struct {
	Fx, Fy float64
	M      float64
	Index  point.Dim2
}
