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
	res := ParseCommand("L12")
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

func TestRun(t *testing.T) {
	got := Run([]Command{
		{Left, 68},
		{Left, 30},
		{Right, 48},
		{Left, 5},
		{Right, 60},
		{Left, 55},
		{Left, 1},
		{Left, 99},
		{Right, 14},
		{Left, 82},
	})
	want := 3
	if got != want {
		t.Fail()
	}
}

func TestRunWithInput(t *testing.T) {
	res := Run()
	t.Log(res)
}