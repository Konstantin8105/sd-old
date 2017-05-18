package coordinate

// PointIndex - alias
type PointIndex int

// Coordinate - position of point
// Base unit for coordinates - meter
type Coordinate struct {
	X, Y, Z float64
	Index   PointIndex
}
