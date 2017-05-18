package model

import (
	"fmt"

	matrix "github.com/Konstantin8105/GoFea/Matrix"
)

// Solve - solving finite element
func (m *Model) Solve() (err error) {

	//TODO: add sort of beam without lost index
	//TODO: avoid absolute free points
	//TODO: avoid mgic number 6 and 12

	// global matrix of stiffiner
	buffer := matrix.NewSquareMatrix(12)
	bufferTr := matrix.NewSquareMatrix(12)
	globalK := matrix.NewSquareMatrix(len(m.coordinates) * 6)
	var fe FiniteElement3Dbeam
	for _, beam := range m.beams {

		material, err := m.GetMaterial(beam.Index)
		if err != nil {
			return fmt.Errorf("Cannot found material for beam #%v. Error = %v", beam.Index, err)
		}
		shape, err := m.GetShape(beam.Index)
		if err != nil {
			return fmt.Errorf("Cannot found shape for beam #%v. Error = %v", beam.Index, err)
		}
		coord, err := m.GetCoordinate(beam.Index)
		if err != nil {
			return fmt.Errorf("Cannot calculate lenght for beam #%v. Error = %v", beam.Index, err)
		}
		fe = FiniteElement3Dbeam{
			material:    material,
			shape:       shape,
			coordinates: coord,
		}
		err = fe.getStiffinerK(&buffer)
		if err != nil {
			return err
		}
		err = fe.getCoordinateTransformation(&bufferTr)
		if err != nil {
			return err
		}

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
	}

	fmt.Println("Global = ", globalK)
	return nil
}
