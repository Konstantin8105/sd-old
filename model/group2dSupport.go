package model

import (
	"github.com/Konstantin8105/GoFea/point"
	"github.com/Konstantin8105/GoFea/support"
)

type supportGroup2d struct {
	support      support.Dim2
	pointIndexes []point.Index
}
