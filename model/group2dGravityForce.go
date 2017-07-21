package model

import (
	"github.com/Konstantin8105/GoFea/element"
	"github.com/Konstantin8105/GoFea/force"
)

type gravityForce2d struct {
	gravityForce force.GravityDim2
	beamIndexes  []element.ElementIndex
}
