package model

import (
	"github.com/Konstantin8105/GoFea/output/displacement"
	"github.com/Konstantin8105/GoFea/output/forceLocal"
	"github.com/Konstantin8105/GoFea/output/reaction"
)

type forceCase2d struct {
	indexCase int

	// input data
	gravityForces []gravityForce2d
	nodeForces    []nodeForce2d

	// statis property
	static staticTypes

	// dynamic property
	dynamicType  dynamicTypes
	dynamicValue []float64

	// output data
	globalDisplacements []displacement.PointDim2
	localDisplacement   []displacement.BeamDim2
	localForces         []forceLocal.BeamDim2
	reactions           []reaction.Dim2
}

type staticTypes bool

const (
	linear staticTypes = false
	nolinear
)

type dynamicTypes int

const (
	none dynamicTypes = iota
	naturalFrequency
	bucklingFactors
)

/*
func zeroCopy(f forceCase2d) (result forceCase2d) {
	result.indexCase = f.indexCase
	result.gravityForces = make([]gravityForce2d, len(f.gravityForces))
	result.nodeForce2d = make([]nodeForce2d, len(f.nodeForces))
	return
}

func delta(a, b forceCase2d) (d float64) {
	for i := range a.gravityForces {
		d = math.Max(d, math.Abs(a.gravityForces[i]-b.gravityForces[i])/math.Max(math.Abs(a.gravityForces[i]), math.Abs(b.gravityForces[i])))
	}
	for i := range a.nodeForces {
		d = math.Max(d, math.Abs(a.nodeForces[i]-b.nodeForces[i])/math.Max(math.Abs(a.nodeForces[i]), math.Abs(b.nodeForces[i])))
	}
	return
}
*/
