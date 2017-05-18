package support

// Dim3 - support of point
type Dim3 struct {
	Dx, Dy, Dz bool // false - free, true - fixed
	Mx, My, Mz bool // false - free, true - fixed
}

// FixedDim3 - fixed support in 3D interpratation
func FixedDim3() (s Dim3) {
	return Dim3{true, true, true, true, true, true}
}
