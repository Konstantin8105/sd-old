package utils_test

import (
	"math"
	"testing"

	"github.com/Konstantin8105/GoFea/input/point"
	"github.com/Konstantin8105/GoFea/utils"
)

func TestLength2DZero(t *testing.T) {
	p0 := point.Dim2{
		X: 0.0,
		Y: 0.0,
	}
	p1 := point.Dim2{
		X: 0.0,
		Y: 0.0,
	}
	if utils.LengthDim2(p0, p1) > 0.0 {
		t.Errorf("Wrong zero length test")
	}
}

func TestLength2DOne1(t *testing.T) {
	p0 := point.Dim2{
		X: 0.0,
		Y: 0.0,
	}
	p1 := point.Dim2{
		X: 1.0,
		Y: 0.0,
	}
	if utils.LengthDim2(p0, p1) != 1.0 {
		t.Errorf("Wrong test with length 1")
	}
}

func TestLength2DOne2(t *testing.T) {
	p0 := point.Dim2{
		X: 0.0,
		Y: 1.0,
	}
	p1 := point.Dim2{
		X: 0.0,
		Y: 0.0,
	}
	if utils.LengthDim2(p0, p1) != 1.0 {
		t.Errorf("Wrong test with length 1")
	}
}

func TestLength2DOne3(t *testing.T) {
	p0 := point.Dim2{
		X: 0.0,
		Y: 0.0,
	}
	p1 := point.Dim2{
		X: 1.0,
		Y: 1.0,
	}
	if math.Abs(utils.LengthDim2(p0, p1)-math.Sqrt(2.0)) > 1e-7 {
		t.Errorf("Wrong test with length 1")
	}
}
