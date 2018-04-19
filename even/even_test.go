package even

import "testing"

func TestEven(t *testing.T) {
	if !Even(8){
		t.Error("8 should be even");
		t.Fail();
	}
	if Even(9) {
		t.Error("9 should not be even");
		t.Fail();
	}

}
