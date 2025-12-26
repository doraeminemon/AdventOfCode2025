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
	res := ReadCommand("L12")
	if res.direction != Left {
		t.Log("Wrong direction")
	}
	if res.increment != 12 {
		t.Log("Wrong increment")
	}
}