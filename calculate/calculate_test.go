package calculate

import "testing"

func TestGetResults(t *testing.T) {

	assertCorrectMapValues := func(t testing.TB, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	}

	double := [][]interface{}{
		{"date1", "0.217"},
		{"date2", "0.201"},
		{"date3", "0.199"},
		{"date4", "0.223"},
		{"date5", "0.216"},
	}

	t.Run("highest exchange rate", func(t *testing.T) {
		got := GetResults(double)["high"]
		want := 0.223
		assertCorrectMapValues(t, got, want)
	})

	t.Run("lowest exchange rate", func(t *testing.T) {
		got := GetResults(double)["low"]
		want := 0.199
		assertCorrectMapValues(t, got, want)
	})

	t.Run("average exchange rate", func(t *testing.T) {
		got := GetResults(double)["average"]
		want := (0.217 + 0.201 + 0.199 + 0.223 + 0.216) / 5
		assertCorrectMapValues(t, got, want)
	})

}
