package model

import (
	"github.com/Konstantin8105/GoFea/input/element"
	"github.com/Konstantin8105/GoFea/input/material"
)

type materialLinearGroup struct {
	material    material.Linear
	beamIndexes []element.Index
}
