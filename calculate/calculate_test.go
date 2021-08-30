package calculate

import "testing"

func TestGetResults(t *testing.T) {

	assertCorrectMapValues := func(t testing.TB, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	}

	double1 := [][]interface{}{
		{"date1", "0.217"},
		{"date2", "0.201"},
		{"date3", "0.199"},
		{"date4", "0.223"},
		{"date5", "0.216"},
	}

	double2 := [][]interface{}{
		{"date1", "0.219"},
		{"date2", "0.222"},
		{"date3", "0.216"},
		{"date4", "0.222"},
		{"date5", "0.216"},
		{"date6", "0.219"},
	}

	t.Run("highest exchange rate for slice of distinct numbers", func(t *testing.T) {
		got := GetResults(double1).High
		want := 0.223
		assertCorrectMapValues(t, got, want)
	})

	t.Run("highest exchange rate with repeated numbers", func(t *testing.T) {
		got := GetResults(double2).High
		want := 0.222
		assertCorrectMapValues(t, got, want)
	})

	t.Run("lowest exchange rate for slice of distinct numbers", func(t *testing.T) {
		got := GetResults(double1).Low
		want := 0.199
		assertCorrectMapValues(t, got, want)
	})

	t.Run("lowest exchange rate with repeated numbers", func(t *testing.T) {
		got := GetResults(double2).Low
		want := 0.216
		assertCorrectMapValues(t, got, want)
	})

	t.Run("average exchange rate for slice of distinct numbers", func(t *testing.T) {
		got := GetResults(double1).Average
		want := (0.217 + 0.201 + 0.199 + 0.223 + 0.216) / 5
		assertCorrectMapValues(t, got, want)
	})

	t.Run("average exchange rate with repeated numbers", func(t *testing.T) {
		got := GetResults(double2).Average
		want := (0.219 + 0.222 + 0.216 + 0.222 + 0.216 + 0.219) / 6
		assertCorrectMapValues(t, got, want)
	})

	t.Run("median exchange rate for slice of distinct numbers", func(t *testing.T) {
		got := GetResults(double1).Median
		want := 0.216
		assertCorrectMapValues(t, got, want)
	})
}
