package model

import (
	"github.com/Konstantin8105/GoFea/input/force"
	"github.com/Konstantin8105/GoFea/input/point"
)

type nodeForce2d struct {
	nodeForce  force.NodeDim2
	pointIndex point.Index
}
