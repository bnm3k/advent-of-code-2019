package day03

import (
	"reflect"
	"testing"
)

func TestParseSingleDirection(t *testing.T) {

	assertCorrectDirection := func(t *testing.T, got, want direction) {
		t.Helper()
		if got.towards != want.towards || got.steps != want.steps {
			t.Errorf("Got %v, want %v\n", got, want)
		}
	}

	assertError := func(t *testing.T, err error) {
		t.Helper()
		if err == nil {
			t.Errorf("Got %v. Want non-nil error\n", err)
		}
	}

	t.Run("test parse single direction", func(t *testing.T) {
		want := direction{right, 9090}
		got, _ := parseDirection("R9090")
		assertCorrectDirection(t, got, want)
	})

	t.Run("test parse single direction with invalid dir", func(t *testing.T) {
		_, err := parseDirection("X9090")
		assertError(t, err)
	})

	t.Run("test parse single direction with invalid num string", func(t *testing.T) {
		_, err := parseDirection("R90,")
		assertError(t, err)
	})

}

func TestParseMultipleDirections(t *testing.T) {
	assertNoError := func(t *testing.T, err error) {
		t.Helper()
		if err != nil {
			t.Errorf("Got %v. Want nil\n", err)
		}
	}

	assertCorrectDirections := func(t *testing.T, got, want []direction) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("\ngot %v \nwant %v\n", got, want)
		}
	}
	dirStr := "R8,U5,L5,D3"
	want := []direction{
		direction{right, 8},
		direction{up, 5},
		direction{left, 5},
		direction{down, 3}}
	got, err := parseMultipleDirections(dirStr)
	assertNoError(t, err)
	assertCorrectDirections(t, got, want)
}

func TestGetMinDistance(t *testing.T) {
	cases := []struct {
		wire1        string
		wire2        string
		cboard       circuitBoard
		expectedDist int
	}{
		{
			"R8,U5,L5,D3",
			"U7,R6,D4,L4",
			newCircuitBoard(50, 50),
			6,
		},
		{
			"R75,D30,R83,U83,L12,D49,R71,U7,L72",
			"U62,R66,U55,R34,D71,R55,D58,R83",
			newCircuitBoard(501, 501),
			159,
		},
		{
			"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			newCircuitBoard(500, 500),
			135,
		},
	}
	assertCorrectDistance := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("\ngot distance %d want %d\n", got, want)
		}
	}
	t.Run("test intCode evaluator", func(t *testing.T) {
		for _, c := range cases {
			got, _ := c.cboard.getClosestDistance(c.wire1, c.wire2)
			assertCorrectDistance(t, got, c.expectedDist)
		}
	})
}
