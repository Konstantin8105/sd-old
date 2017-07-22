package support

// Dim2 - support in 2d
type Dim2 struct {
	Dx, Dy Type
	M      Type
}

// FixedDim2 - fixed support in 2D interpratation
func FixedDim2() (s Dim2) {
	return Dim2{Fix, Fix, Fix}
}
