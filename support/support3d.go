package support

// Dim3 - support of point
type Dim3 struct {
	Dx, Dy, Dz Type
	Mx, My, Mz Type
}

// FixedDim3 - fixed support in 3D interpratation
func FixedDim3() (s Dim3) {
	return Dim3{
		Fix, Fix, Fix,
		Fix, Fix, Fix,
	}
}
