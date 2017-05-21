package model

import (
	"fmt"

	"github.com/Konstantin8105/GoFea/dof"
	"github.com/Konstantin8105/GoFea/finiteElement"
	"github.com/Konstantin8105/GoFea/linearAlgebra"
)

// Solve - solving finite element
func (m *Dim2) Solve() (err error) {

	//TODO: add sort of beam without lost index
	//TODO: avoid absolute free points
	//TODO: avoid mgic number 6 and 12

	// global matrix of stiffiner
	buffer := linearAlgebra.NewSquareMatrix(12)
	bufferTr := linearAlgebra.NewSquareMatrix(12)
	bufferM := linearAlgebra.NewSquareMatrix(12)
	dof := dof.NewBeam(m.beams, dof.Dim2d)
	//globalK := linearAlgebra.NewSquareMatrix(len(m.points) * 6)
	for _, beam := range m.beams {

		material, err := m.getMaterial(beam.Index)
		if err != nil {
			return fmt.Errorf("Cannot found material for beam #%v. Error = %v", beam.Index, err)
		}
		shape, err := m.getShape(beam.Index)
		if err != nil {
			return fmt.Errorf("Cannot found shape for beam #%v. Error = %v", beam.Index, err)
		}
		coord, err := m.getCoordinate(beam.Index)
		if err != nil {
			return fmt.Errorf("Cannot calculate lenght for beam #%v. Error = %v", beam.Index, err)
		}
		if m.isTruss(beam.Index) {
			tr := finiteElement.TrussDim2{
				Material: material,
				Shape:    shape,
				Points:   coord,
			}
			err = tr.GetStiffinerK(&buffer)
			if err != nil {
				return err
			}

			//fmt.Println("K", beam.Index, " =\n", buffer)

			err = tr.GetCoordinateTransformation(&bufferTr)
			if err != nil {
				return err
			}

			//fmt.Println("T", beam.Index, " =\n", bufferTr)

			Kor := buffer.MultiplyTtKT(bufferTr, &bufferM)

			fmt.Println("Kor", beam.Index, " =\n", Kor)

			tr.GetDoF(&dof)

			fmt.Println("DoF = ", tr.Axe)

		}
		/* else {
			fe := finiteElement.BeamDim2{
				Material: material,
				Shape:    shape,
				Points:   coord,
			}
			err = fe.GetStiffinerK(&buffer)
			if err != nil {
				return err
			}
		}*/

		/*

			fmt.Println(buffer)
			inx0 := (int(coord[0].Index) - 1) * 6
			for i := 0; i < 6; i++ {
				for j := 0; j < 6; j++ {
					globalK.Set(inx0+i, inx0+j, globalK.Get(inx0+i, inx0+j)+buffer.Get(i, j))
				}
			}
			inx1 := (int(coord[1].Index) - 1) * 6
			for i := 0; i < 6; i++ {
				for j := 0; j < 6; j++ {
					globalK.Set(inx1+i, inx1+j, globalK.Get(inx1+i, inx1+j)+buffer.Get(i+6, j+6))
				}
			}
			for i := 0; i < 6; i++ {
				for j := 0; j < 6; j++ {
					globalK.Set(inx0+i, inx1+j, globalK.Get(inx0+i, inx1+j)+buffer.Get(i+6, j))
				}
			}
			for i := 0; i < 6; i++ {
				for j := 0; j < 6; j++ {
					globalK.Set(inx1+i, inx0+j, globalK.Get(inx1+i, inx0+j)+buffer.Get(i, j+6))
				}
			}
		*/
	}

	//fmt.Println("Global = ", globalK)
	return nil
}
