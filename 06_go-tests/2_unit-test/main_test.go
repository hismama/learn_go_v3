package main

import "testing"

func TestAdd(t *testing.T) {
	total := add(4, 11)
	if total != 15 {
		t.Errorf("Add(4, 11) = %d; want 15", total)
	}
}

func TestSubtract(t *testing.T) {
	total := subtract(17, 11)
	if total != 6 {
		t.Errorf("Subtract(17, 11) = %d; want 6", total)
	}
}

func TestDoMath1(t *testing.T) {
	total := doMath(4, 11, add)
	if total != 15 {
		t.Errorf("doMath(4, 11, add) = %d; want 15", total)
	}
}

func TestDoMath2(t *testing.T) {
	total := doMath(4, 11, subtract)
	if total != -7 {
		t.Errorf("doMath(4, 11, subtract) = %d; want -7", total)
	}
}
