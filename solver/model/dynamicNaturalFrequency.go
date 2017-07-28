package model

/*
	//TODO: can calculated in parallel local force

	// Generate global mass matrix [Mo]
	n := stiffinerKGlobal.GetRowSize()
	massGlobal := m.convertFromLocalToGlobalSystem(&degreeGlobal, &dofSystem, &mapIndex, finiteElement.GetGlobalMass)
	// m.convertFromLocalToGlobalSystem(&degreeGlobal, &dofSystem, &mapIndex, finiteElement.GetGlobalMass)
	//  linAlg.NewMatrix64bySize(n, n)

	// TODO: Add to matrix mass the nodal mass
	for _, node := range forceCase.nodeForces {
		for _, inx := range node.pointIndexes {
			d := dofSystem.GetDoF(inx)
			if node.nodeForce.Fx != 0.0 {
				h, err := mapIndex.GetByAxe(d[0])
				if err == nil {
					massGlobal.Set(h, h, massGlobal.Get(h, h)+math.Abs(node.nodeForce.Fx))
				}
			}
			if node.nodeForce.Fy != 0.0 {
				h, err := mapIndex.GetByAxe(d[1])
				if err == nil {
					massGlobal.Set(h, h, massGlobal.Get(h, h)+math.Abs(node.nodeForce.Fy))
				}
			}
			// TODO: Moment haven`t mass ???
			// TODO: Check
			//if node.nodeForce.M != 0.0 {
			//	h, err := mapIndex.GetByAxe(d[2])
			//	if err == nil {
			//		massGlobal.Set(h, h, massGlobal.Get(h, h)+math.Abs(node.nodeForce.M))
			//		fmt.Println("Add M to mass")
			//	}
			//}
		}
	}

	//TODO: CHECKUING GRAVITY TO MATRIX MASS
	for i := 0; i < massGlobal.GetRowSize(); i++ {
		for j := 0; j < massGlobal.GetColumnSize(); j++ {
			massGlobal.Set(i, j, massGlobal.Get(i, j)/9.806)
		}
	}
	// TODO: ADD to mass WITH OR WITOUT SELFWEIGHT

	// Calculate matrix [H] = [Ko]^-1 * [Mo]
	if stiffinerKGlobal.GetRowSize() != stiffinerKGlobal.GetColumnSize() {
		panic("Not correct size of global stiffiner matrix")
	}
	//fmt.Println("GlobalMass = ", massGlobal)
	Ho := matrix.NewMatrix64bySize(n, n)
	buffer := matrix.NewMatrix64bySize(n, 1)
	for i := 0; i < n; i++ {
		// Create vertical vector from [Mo]
		for j := 0; j < n; j++ {
			buffer.Set(j, 0, massGlobal.Get(j, i))
		}
		// Calculation
		result := lu.Solve(buffer)
		// Add vector to [Ho]
		for j := 0; j < n; j++ {
			Ho.Set(j, i, result.Get(j, 0))
		}
	}
	//fmt.Println("[Ho] = ", Ho)
	{
		// TODO: check
		// Remove zero rows and columns
		var removePosition []int
		// TODO: len --> to matrix length
		// TODO: at the first check diagonal element
		for i := 0; i < Ho.GetRowSize(); i++ {
			found := false
			for j := 0; j < Ho.GetRowSize(); j++ {
				if Ho.Get(i, j) != 0.0 {
					found = true
					break
				}
			}
			if found {
				continue
			}
			removePosition = append(removePosition, i)
		}
		Ho.RemoveRowAndColumn(removePosition...)
	}

	// Calculation of natural frequency
	eigen := solver.NewEigen(Ho)
	//fmt.Println("lambda       = ", eigen.GetRealEigenvalues())
	//fmt.Println("eigenvectors = ", eigen.GetV())
	//fmt.Println("getD = ", eigen.GetD())

	// TODO: fix for avoid strange frequency some is too small or too big
	value := eigen.GetRealEigenvalues()
	for _, v := range value {
		freq := math.Sqrt(1.0/v) / 2.0 / math.Pi
		fmt.Printf("f = %.5v Hz\n", freq)
		_ = freq
	}
	// TODO: need add modal mass values for natural frquency calculation

*/
