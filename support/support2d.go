package support

type Dim2 struct {
	Dx, Dy bool // false - free, true - fixed
	M      bool // false - free, true - fixed
}

// FixedDim2 - fixed support in 2D interpratation
func FixedDim2() (s Dim2) {
	return Dim2{true, true, true}
}
