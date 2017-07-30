package finiteElement_test

import (
	"fmt"
	"testing"

	"github.com/Konstantin8105/GoFea/solver/dof"
	"github.com/Konstantin8105/GoFea/solver/finiteElement"
	"github.com/Konstantin8105/GoLinAlg/matrix"
)

func TestRemovePanic(t *testing.T) {
	tests := []struct {
		m matrix.T64
		a []dof.AxeNumber
	}{
		{
			m: matrix.NewMatrix64bySize(1, 1),
			a: []dof.AxeNumber{1, 2, 3},
		},
		{
			m: matrix.NewMatrix64bySize(5, 5),
			a: []dof.AxeNumber{1, 2, 3},
		},
	}
	for index, test := range tests {
		t.Run(fmt.Sprintf("Remove-%v", index), func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("The code did not panic")
				}
			}()
			finiteElement.RemoveZeros(&test.m, &test.a)
		})
	}
}
