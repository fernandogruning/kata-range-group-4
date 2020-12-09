package main

import "testing"

func TestValidConstructor(t *testing.T) {
	rRange := Range{"[2,5)"}
	r := rRange.ToRealRange()
	got := r
	want := RealRange{2, true, 5, false}

	if got != want {
		t.Errorf("got %v wannt %v", got, want)
	}
}
