package force

// NodeDim2 - force on node in 2D interpretation
// Unit - N
type NodeDim2 struct {
	Fx, Fy float64
	M      float64
}

// Plus - add(summary) load
func (n *NodeDim2) Plus(a NodeDim2) {
	n.Fx += a.Fx
	n.Fy += a.Fy
	n.M += a.M
}
