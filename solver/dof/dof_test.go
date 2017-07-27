package dof_test

import (
	"testing"

	"github.com/Konstantin8105/GoFea/solver/dof"
)

func TestFound(t *testing.T) {
	var d dof.DoF
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	_ = d.GetDoF(1)
}
