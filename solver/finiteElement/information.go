package finiteElement

// Information - type of information for
// global stiffiner matrix and
// slise of dof
type Information bool

// Information constants
const (
	WithoutZeroStiffiner Information = false
	FullInformation                  = true
)
