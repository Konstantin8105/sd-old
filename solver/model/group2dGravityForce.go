package model

import (
	"github.com/Konstantin8105/GoFea/input/element"
	"github.com/Konstantin8105/GoFea/input/force"
)

type gravityForce2d struct {
	gravityForce force.GravityDim2
	beamIndexes  []element.ElementIndex
}
