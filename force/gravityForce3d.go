package force

// GravityDim3 - gravity acceleration for self-weight
// Unit - m/s^2
// For example, for case with gravity by Z - Gx=0,Gy=0,Gz=-9.8
type GravityDim3 struct {
	Gx, Gy, Gz float64
}
