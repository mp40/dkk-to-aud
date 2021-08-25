package calculate

import (
	"strconv"
)

func GetResults(values [][]interface{}) map[string]float64 {

	high := 0.0
	low := 666.0
	sum := 0.0

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
	}

	results := map[string]float64{
		"high":    high,
		"low":     low,
		"average": sum / float64(len(values)),
	}

	return results
}
