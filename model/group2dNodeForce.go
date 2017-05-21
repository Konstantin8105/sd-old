package model

import (
	"github.com/Konstantin8105/GoFea/force"
	"github.com/Konstantin8105/GoFea/point"
)

type nodeForce2d struct {
	nodeForce    force.NodeDim2
	pointIndexes []point.Index
}
