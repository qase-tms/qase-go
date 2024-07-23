package _go

import "testing"

func Test_SumSuccess(t *testing.T) {
	result := Sum(1, 2)
	if result != 3 {
		t.Errorf("Sum(1, 2) = %d; want 3", result)
	}
}

func Test_SumFailed(t *testing.T) {
	result := Sum(1, 2)
	if result != 4 {
		t.Errorf("Sum(1, 2) = %d; want 4", result)
	}
}
