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
