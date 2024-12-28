package main

import "testing"

func TestAdd(t *testing.T) {
	total := add(4, 11)
	if total != 15 {
		t.Errorf("Add(4, 11) = %d; want 15", total)
	}
}
