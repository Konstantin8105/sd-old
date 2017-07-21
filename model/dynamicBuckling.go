package model

/*

	// Linear buckling
	//potentialGlobal := m.convertFromLocalToGlobalSystem(&degreeGlobal, &dofSystem, &mapIndex, finiteElement.GetGlobalPotential)
	potentialGlobal := matrix.NewMatrix64bySize(stiffinerKGlobal.GetRowSize(), stiffinerKGlobal.GetColumnSize())
	for _, beam := range m.beams {
		fe := m.getBeamFiniteElement(beam.Index)

		//klocal,
		_, degreeLocal := finiteElement.GetStiffinerGlobalK(fe, &dofSystem, finiteElement.FullInformation)
		globalDisplacement := make([]float64, len(degreeLocal))
		// if not found in global displacement, then it is a pinned
		// in local stiffiner matrix - than row and column is zero
		// for avoid collisian - we put a zero
		for i := 0; i < len(globalDisplacement); i++ {
			for j := 0; j < len(degreeGlobal); j++ {
				if degreeLocal[i] == degreeGlobal[j] {
					globalDisplacement[i] = x.Get(j, 0)
					break
				}
			}
		}

		t := matrix.NewMatrix64bySize(10, 10)
		fe.GetCoordinateTransformation(&t)

		// Zo = T_t * Z
		var localDisplacement []float64
		for i := 0; i < t.GetRowSize(); i++ {
			sum := 0.0
			for j := 0; j < t.GetColumnSize(); j++ {
				sum += t.Get(i, j) * globalDisplacement[j]
			}
			localDisplacement = append(localDisplacement, sum)
		}

		kk := matrix.NewMatrix64bySize(10, 10)
		fe.GetStiffinerK(&kk)

		var localForce []float64
		for i := 0; i < kk.GetRowSize(); i++ {
			sum := 0.0
			for j := 0; j < kk.GetRowSize(); j++ {
				sum += kk.Get(i, j) * localDisplacement[j]
			}
			localForce = append(localForce, sum)
		}

		//fmt.Println("local Force = ", localForce)
		if localForce[0] > 0.0 && localForce[3] < 0.0 {
			// TODO : it is not correct , because uniform load can change
			//fmt.Println("Compress")
		} else {
			// TODO: testing
			localForce[0] = 0.0
		}

		grLocal := matrix.NewMatrix64bySize(6, 6)
		fe.GetPotentialGr(&grLocal, localForce[0])

		// Add local stiffiner matrix to global matrix
		for i := 0; i < len(degreeLocal); i++ {
			g, err := mapIndex.GetByAxe(degreeLocal[i])
			if err != nil {
				continue
			}
			for j := 0; j < len(degreeLocal); j++ {
				h, err := mapIndex.GetByAxe(degreeLocal[j])
				if err != nil {
					continue
				}
				potentialGlobal.Set(g, h, potentialGlobal.Get(g, h)+grLocal.Get(i, j))
			}
		}
	}

	//fmt.Println("PotentialGlobal = ", potentialGlobal)
	HoPotential := matrix.NewMatrix64bySize(n, n)
	bufferPotential := matrix.NewMatrix64bySize(n, 1)
	//fmt.Printf("lu = %#v\n", lu)
	for i := 0; i < n; i++ {
		// Create vertical vector from [Mo]
		for j := 0; j < n; j++ {
			bufferPotential.Set(j, 0, potentialGlobal.Get(j, i))
		}
		// Calculation
		result := lu.Solve(bufferPotential)
		// Add vector to [Ho]
		for j := 0; j < n; j++ {
			HoPotential.Set(j, i, result.Get(j, 0))
		}
	}
	//fmt.Println("[HoPotential] = ", HoPotential)
	{
		// TODO: check
		// Remove zero rows and columns
		var removePosition []int
		// TODO: len --> to matrix lenght
		// TODO: at the first check diagonal element
		for i := 0; i < HoPotential.GetRowSize(); i++ {
			found := false
			for j := 0; j < HoPotential.GetRowSize(); j++ {
				if HoPotential.Get(i, j) != 0.0 {
					found = true
					break
				}
			}
			if found {
				continue
			}
			removePosition = append(removePosition, i)
		}
		HoPotential.RemoveRowAndColumn(removePosition...)
	}
	// Calculation of
	eigenPotential := solver.NewEigen(HoPotential)
	//fmt.Println("lambda       = ", eigenPotential.GetRealEigenvalues())
	//fmt.Println("lambda Re    = ", eigenPotential.GetImagEigenvalues())
	//fmt.Println("eigenvectors = ", eigenPotential.GetV())
	//fmt.Println("getD = ", eigenPotential.GetD())

	// TODO: Remove strange results
	valueP := eigenPotential.GetRealEigenvalues()
	fmt.Println("Linear buckling loads:")
	for _, v := range valueP {
		fmt.Printf("P = %.5v\n", 1.0/v)
	}

	///  BUckling iteration solving
	// [K]  = stiffinerKGlobal
	// [Kg] = potentialGlobal
	//	for iter := 0; iter < 1000; iter++ {

	//}

	// Nolinear buckling calculation
	// algorithm Newton-Raphfon
	// для дальнейшего развития
	// необходимо рекурсивно вызывать эту
	// функцию Solve
	type step struct {
		forces       forceCase2d
		displacement matrix.T64
	}
	type iteration struct {
		s      step
		result resultNolinearBuckling
	}
	/*
		loadsOld := zeroCopy(m.forceCases[caseNumber])
		displacementOld := zeroDisplacement(...)
		resultOld := converge
		displacementOld, resultOld = calculate(loadsOld)

		loadsNew := m.forceCases[caseNumber]
		var displacementNew Matrix64
		var resultNew  resultNolinearBuckling
		displacementNew, resultNew = calculate(loadsNew)

		for {
			if resultNew == diverge{
				break
			}
			loadsOld = loadsNew
			resultOld = resultNew
			loadsNew = multiply(2.0, loadsNew)
		}

		eps := 0.01
		for {

			if deltaDisp(displacementNew, displacementOld) <= eps && deltaLoads(loadsNew, loadsOld) <= eps && resultOld == converge && resultNew == diverge {
				break
			}
			loadAverage := average(loadsNew, loadsOld)
			switch resultAverage{
			case converge:
				loadsOld = loadsAverage
			case diverge:
				loadsNew = loadsAverage
			}
		}
*/
