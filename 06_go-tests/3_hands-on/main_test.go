package main

import "testing"

func TestSRandom(t *testing.T) {
	got := sRandom("I've been.")
	expected := "What a machine I've been."
	if got != expected {
		t.Errorf("\nexpected:\t'%s'\ngot:\t\t'%s'", expected, got)
	}
}
