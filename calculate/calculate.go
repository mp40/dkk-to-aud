package calculate

import (
	"encoding/json"
	"log"
	"math"
	"sort"
	"strconv"
	"time"

	"github.com/mp40/dkk-to-aud/cache"
)

type Data struct {
	High    float64
	Low     float64
	Latest  float64
	Average float64
	Median  float64
}

func GetMedian(sorted []float64, length int) float64 {
	m := length / 2

	if length%2 == 0 {
		return (sorted[m] + sorted[m-1]) / 2
	}

	return sorted[m]
}

func cacheData(results *Data) {
	timestamp := time.Now().UTC()
	key := timestamp.Format("2006-01-02")

	value, marshalErr := json.Marshal(results)

	cache := &cache.Cache

	if marshalErr == nil {
		_, err := cache.Set(key, string(value))
		if err != nil {
			log.Println(err)
		}
	}
}

func GetResults(values [][]interface{}) *Data {

	high := 0.0
	low := 666.0
	sum := 0.0

	l := len(values)

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

	latest := s[len(s)-1]
	sort.Float64s(s)

	results := Data{
		High:    high,
		Low:     low,
		Latest:  latest,
		Average: math.Round((sum/float64(l))*10000000) / 10000000,
		Median:  GetMedian(s, l),
	}

	cacheData(&results)

	return &results
}
