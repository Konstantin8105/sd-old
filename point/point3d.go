package point

// Index - alias
type Index int

// Dim3 - position of point in 3d (dimentions)
// Base unit for coordinates - meter
type Dim3 struct {
	X, Y, Z float64
	Index   Index
}
