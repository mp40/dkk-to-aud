package calculate

import "testing"

func TestGetResults(t *testing.T) {

	double := [][]interface{}{
		{"date1", "0.217"},
		{"date2", "0.201"},
		{"date3", "0.199"},
		{"date4", "0.223"},
		{"date5", "0.216"},
	}

	got := GetResults(double)["high"]

	want := 0.223

	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}
