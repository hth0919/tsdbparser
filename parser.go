package tsdbParser

import "encoding/json"

func JsonUnmarshaller (jsonfile []byte) *Metric {
	mt := &Metric{
		MetricName: "",
		Tags: map[string]string{},
		Dps: map[int64]int{},
	}
	err :=json.Unmarshal(jsonfile, mt)
	if err != nil {
		panic(err)
	}

	return mt
}

func GetLast(input *Metric) int {
	var last int64
	last = 0
	for k := range input.Dps {
		if k>last {
			last = k
		}
	}

	return input.Dps[last]
}