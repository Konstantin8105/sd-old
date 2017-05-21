package model

import (
	"github.com/Konstantin8105/GoFea/element"
	"github.com/Konstantin8105/GoFea/material"
)

type materialLinearGroup struct {
	material    material.Linear
	beamIndexes []element.BeamIndex
}
