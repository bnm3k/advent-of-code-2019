package day01

import "testing"

func TestTask1(t *testing.T) {
	casesTask2 := []struct {
		Weight       int
		ExpectedFuel int
	}{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}

	assertCorrectFuel := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("test get fuel for task 2", func(t *testing.T) {
		for _, test := range casesTask2 {
			got := getFuelTask2(test.Weight)
			assertCorrectFuel(t, got, test.ExpectedFuel)
		}
	})

}
