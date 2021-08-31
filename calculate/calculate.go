package calculate

import (
	"sort"
	"strconv"
)

type Data struct {
	High    float64
	Low     float64
	Average float64
	Median  float64
}

func GetMedian(sorted []float64, len int) float64 {
	m := len / 2

	if len%2 == 0 {
		return (sorted[m] + sorted[m-1]) / 2
	}

	return sorted[m]
}

func GetResults(values [][]interface{}) *Data {

	high := 0.0
	low := 666.0
	sum := 0.0

	len := len(values)

	s := make([]float64, 0)

	for _, v := range values {

		var f float64
		var err error

		if str, ok := v[1].(string); ok {
			f, err = strconv.ParseFloat(str, 64)

			if err != nil {
				panic(err)
			}
		} else {
			f = 0.0
		}

		if f > high {
			high = f
		}

		if f < low {
			low = f
		}

		sum += f

		s = append(s, f)
	}

	sort.Float64s(s)

	results := Data{
		High:    high,
		Low:     low,
		Average: sum / float64(len),
		Median:  GetMedian(s, len),
	}

	return &results
}
