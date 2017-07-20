package main_test

/*
func ExampleFrame3DD() {

	var model model.Dim2

	// input data of exB

	coordinates := []point.Dim3{
		point.Dim3{Index: 1, X: 0.0, Y: 0.0, Z: 1.0},
		point.Dim3{Index: 2, X: -1.2, Y: -0.9, Z: 0.0},
		point.Dim3{Index: 3, X: 1.2, Y: -0.9, Z: 0.0},
		point.Dim3{Index: 4, X: 1.2, Y: 0.9, Z: 0.0},
		point.Dim3{Index: 5, X: -1.2, Y: 0.9, Z: 0.0},
	}

	model.AddPoint(coordinates...)

	beams := []element.Beam{
		element.Beam{
			Index:        1,
			PointIndexes: [2]point.Index{2, 1},
		},
		element.Beam{
			Index:        2,
			PointIndexes: [2]point.Index{1, 3},
		},
		element.Beam{
			Index:        3,
			PointIndexes: [2]point.Index{1, 4},
		},
		element.Beam{
			Index:        4,
			PointIndexes: [2]point.Index{5, 1},
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

	steel := material.Linear{
		E:  200000e6,
		G:  79300e6,
		V:  0.3,
		Ro: 78500,
	}

	model.AddMaterial(steel, beamIndexes...)

	supCoordinateIndex := []point.Index{2, 3, 4, 5}

	model.AddSupport(support.FixedDim3(), supCoordinateIndex...)

	gravity := force.GravityDim3{
		Gx: 0.0,
		Gy: 0.0,
		Gz: -9.8033,
	}

	nodeLoad := force.NodeDim3{
		Fx: 100.0,
		Fy: -200.0,
		Fz: -100.0,
	}

	model.AddGravityForce(gravity, 1, 2, 3, 4)
	model.AddNodeForce(nodeLoad, 1)

	//flagShearDeformation := 1     // shear
	//flagGeomerticDeformation := 1 // geom

	//fmt.Printf("Model = %#v\n", model)

	err := model.Solve()
	if err != nil {
		fmt.Println("Error in solving. Err = ", err)
		return
	}
	fmt.Println("OK")
	// Output: OK
}
*/
