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

func TestValidConstructorSCE012(t *testing.T) {
	rRange := Range{"(3,6]"}
	r := rRange.ToRealRange()
	got := r
	want := RealRange{3, false, 6, true}

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

func TestGetAllPointsSCE312(t *testing.T) {
	rRange := Range{"[1,10)"}
	r := rRange.GetAllPoints()
	got := r
	want := "{1,2,3,4,5,6,7,8,9}"

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestContainsRangeSCE411(t *testing.T) {
	rRange := Range{"[2,5)"}

	got := rRange.ContainsRange(Range{"[7,10)"})
	want := "[2,5) doesn't contain [7,10)"

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestContainsRangeSCE412(t *testing.T) {
	rRange := Range{"[2,5)"}

	got := rRange.ContainsRange(Range{"[3,10)"})
	want := "[2,5) doesn't contain [3,10)"

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestContainsRangeSCE413(t *testing.T) {

	rRange := Range{"[2,5)"}

	got := rRange.ContainsRange(Range{"[3,5)"})
	want := "[2,5) contains [3,5)"

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestContainsRangeSCE414(t *testing.T) {

	rRange := Range{"[3,5]"}

	got := rRange.ContainsRange(Range{"[3,5)"})
	want := "[3,5] contains [3,5)"

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
