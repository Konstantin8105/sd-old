package support

// Support - support of point
type Support struct {
	Dx, Dy, Dz bool // false - free, true - fixed
	Mx, My, Mz bool // false - free, true - fixed
}

// FixedSupport - fixed support
func FixedSupport() (s Support) {
	return Support{true, true, true, true, true, true}
}
