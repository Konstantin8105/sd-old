package support_test

import (
	"testing"

	"github.com/Konstantin8105/GoFea/support"
)

func Test2d(t *testing.T) {
	s := support.FixedDim2()
	if s.Dx != support.Fix || s.Dy != support.Fix || s.M != support.Fix {
		t.Errorf("Not correct fixed 2d support")
	}
}

func Test3d(t *testing.T) {
	s := support.FixedDim3()
	if s.Dx != support.Fix || s.Dy != support.Fix || s.Dz != support.Fix || s.Mx != support.Fix || s.My != support.Fix || s.Mz != support.Fix {
		t.Errorf("Not correct fixed 3d support")
	}
}
