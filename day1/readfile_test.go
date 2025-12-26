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

func TestTurnLeftNegative(t *testing.T) {
	got := Turn(50, Command{Left, 68})
	want := ClockPosition(82)
	if (got != want) {
		t.Fail()
	}
}

func TestTurnLefttPositive(t *testing.T) {
	got := Turn(5 , Command{Left, 5})
	want := ClockPosition(0)
	if (got != want) {
		t.Fail()
	}
}

func TestTurnRightPositive(t *testing.T) {
	got := Turn(0, Command{Right, 5})
	want := ClockPosition(5)
	if (got != want) {
		t.Fail()
	}
}

func TestTurnRightNegative(t *testing.T) {
	got := Turn(50, Command{Right, 68})
	want := ClockPosition(18)
	if (got != want) {
		t.Fail()
	}
}