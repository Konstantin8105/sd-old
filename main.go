package main

import (
	"fmt"

	coordinate "github.com/Konstantin8105/GoFea/Coordinate"
	element "github.com/Konstantin8105/GoFea/Element"
	force "github.com/Konstantin8105/GoFea/Force"
	material "github.com/Konstantin8105/GoFea/Material"
	model "github.com/Konstantin8105/GoFea/Model"
	shape "github.com/Konstantin8105/GoFea/Shape"
	support "github.com/Konstantin8105/GoFea/Support"
)

/*

+--------------+
| Element      |
| library      |
+--------------+
     |    |
     |    | Ke
     |    |
+--------------+
| Assembler    |
+--------------+

add finite element - truss, gap, tension, compress element


*/

func main() {

	var model model.Model

	// input data of exB

	coordinates := []coordinate.Coordinate{
		coordinate.Coordinate{Index: 1, X: 0.0, Y: 0.0, Z: 1.0},
		coordinate.Coordinate{Index: 2, X: -1.2, Y: -0.9, Z: 0.0},
		coordinate.Coordinate{Index: 3, X: 1.2, Y: -0.9, Z: 0.0},
		coordinate.Coordinate{Index: 4, X: 1.2, Y: 0.9, Z: 0.0},
		coordinate.Coordinate{Index: 5, X: -1.2, Y: 0.9, Z: 0.0},
	}

	model.AddCoodinates(coordinates...)

	beams := []element.Beam{
		element.Beam{
			Index:       1,
			PointIndexs: [2]coordinate.PointIndex{2, 1},
		},
		element.Beam{
			Index:       2,
			PointIndexs: [2]coordinate.PointIndex{1, 3},
		},
		element.Beam{
			Index:       3,
			PointIndexs: [2]coordinate.PointIndex{1, 4},
		},
		element.Beam{
			Index:       4,
			PointIndexs: [2]coordinate.PointIndex{5, 1},
		},
	}

	model.AddBeams(beams...)

	beamIndexes := []element.BeamIndex{1, 2, 3, 4}

	shape := shape.Shape{
		A:   36.0E-6,
		Asy: 20.0E-6,
		Asz: 20.0E-6,
		Jxx: 1000.0E-12,
		Iyy: 492.0E-12,
		Izz: 491.0E-12,
	}

	model.AddShape(shape, beamIndexes...)

	steel := material.Material{
		E:  200000e6,
		G:  79300e6,
		V:  0.3,
		Ro: 78500,
	}

	model.AddMaterial(steel, beamIndexes...)

	supCoordinateIndex := []coordinate.PointIndex{2, 3, 4, 5}

	model.AddSupport(support.FixedSupport(), supCoordinateIndex...)

	gravity := force.GravityForce{
		Gx: 0.0,
		Gy: 0.0,
		Gz: -9.8033,
	}

	nodeLoad := force.NodeForce{
		Fx: 100.0,
		Fy: -200.0,
		Fz: -100.0,
	}

	model.AddGravityForce(gravity, 1, 2, 3, 4)
	model.AddNodeForce(nodeLoad, 1)

	//flagShearDeformation := 1     // shear
	//flagGeomerticDeformation := 1 // geom

	fmt.Printf("Model = %#v\n", model)

	err := model.Solve()
	if err != nil {
		fmt.Println("Error in solving. Err = ", err)
		return
	}
}
