package tsdbParser

import (
	"encoding/json"
	"io"
	"strconv"
)



func JsonUnmarshaller (jsonfile io.ReadCloser) *Metric {
	mt := &Metric{
		MetricName: "",
		Tags: map[string]string{},
		Dps: map[string]float64{},
	}
	err :=json.NewDecoder(jsonfile).Decode(mt)
	if err != nil {
		panic(err)
	}

	return mt
}

func GetLast(input *Metric) float64 {
	var last int
	last = 0
	for k := range input.Dps {
		i,err := strconv.Atoi(k)
		if err != nil {
			panic(err)
		}
		if i>last {
			last = i
		}
	}

	l := strconv.Itoa(last)

	return input.Dps[l]
}