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
