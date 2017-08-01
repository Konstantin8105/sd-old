package model

/*

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
