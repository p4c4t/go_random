package main

import (
	"testing"
)

func TestSqrt(t *testing.T) {
	got_int := Sqrt(2)
	if got_int != 1 {
		t.Errorf("Sqrt[](2) = %d; want 1", got_int)
	}
	got_complex := Sqrt(0 + 2i)
	if got_complex != 1+1i {
		t.Errorf("Sqrt[](0 + 2i) = %v; want 1", got_complex)
	}
}
