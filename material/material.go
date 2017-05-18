package material

// Linear - linear property of material
type Linear struct {
	E  float64 // Young's modulus, Pa
	G  float64 // Shear modulus, Pa
	V  float64 // Poisson's ratio
	Ro float64 // Density, N/m3
}
