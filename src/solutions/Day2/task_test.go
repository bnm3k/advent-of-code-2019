package main

import (
	"reflect"
	"testing"
)

func TestTask(t *testing.T) {
	cases := []struct {
		InitialState       []int
		FinalExpectedState []int
	}{
		{
			[]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
			[]int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
		},
		{
			[]int{1, 0, 0, 0, 99},
			[]int{2, 0, 0, 0, 99},
		},
		{
			[]int{2, 3, 0, 3, 99},
			[]int{2, 3, 0, 6, 99},
		},
		{
			[]int{2, 4, 4, 5, 99, 0},
			[]int{2, 4, 4, 5, 99, 9801},
		},
		{
			[]int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			[]int{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
	}

	assertCorrectState := func(t *testing.T, initialState, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("\ninitialState %v\ngot %v \nwant %v\n", initialState, got, want)
		}
	}

	t.Run("test intCode evaluator", func(t *testing.T) {
		for _, testCase := range cases {
			got := evaluateIntCode(testCase.InitialState)
			assertCorrectState(t, testCase.InitialState, got, testCase.FinalExpectedState)
		}
	})

	t.Run("test intCode preprocessor", func(t *testing.T) {
		initialState := []int{0, 1, 2, 3, 4}
		got := preprocessIntCode(initialState)
		want := []int{0, 12, 2, 3, 4}
		assertCorrectState(t, initialState, got, want)
	})

}
