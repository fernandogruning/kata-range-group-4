package main

import "testing"

func TestValidConstructorSCE011(t *testing.T) {
	rRange := Range{"[2,5)"}
	r := rRange.ToRealRange()
	got := r
	want := RealRange{2, true, 5, false}

	if got != want {
		t.Errorf("got %v wannt %v", got, want)
	}
}

func TestGetAllPointsSCE311(t *testing.T) {
	rRange := Range{"[2,6)"}
	r := rRange.GetAllPoints()
	got := r
	want := "{2,3,4,5}"

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
