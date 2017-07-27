package element_test

import (
	"testing"

	"github.com/Konstantin8105/GoFea/input/element"
)

func TestBeam(t *testing.T) {
	b := element.NewBeam(0, 0, 0)
	if b.GetAmountPoint() != 2 {
		t.Errorf("Wrong amount points for beam")
	}
}
