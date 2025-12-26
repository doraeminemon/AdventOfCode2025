package day1

import (
	"testing"
)

func TestReadFile(t *testing.T) {
	result := ReadFile()
	got := len(result)
	want := 4046
	if (got != want) {
		t.Log("Unexpected length")
	}
}

func TestReadCommandLeft(t *testing.T) {
	dir, inc := ReadCommand("L12")
	if dir != Left {
		t.Log("Wrong direction")
	}
	if inc != 12 {
		t.Log("Wrong increment")
	}
}