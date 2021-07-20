package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type Data struct {
	Metrics []Metric
}

type Metric struct {
	Value float32 `json:"metricValue"`
	Date  DTime   `json:"dtime"`
}

func LoadDataFromFile(path string) (*Data, error) {
	var data Data

	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &data.Metrics)
	return &data, err
}

func (d *Data) Process() {
	fmt.Println(d)
}

func (d *Data) String() string {
	if len(d.Metrics) == 0 {
		return "" +
			"SamKnows Metric Analyser v1.0.0\n" +
			"===============================\n" +
			"\n" +
			"No metrics processed"
	}

	return fmt.Sprintf(""+
		"SamKnows Metric Analyser v1.0.0\n"+
		"===============================\n"+
		"\n"+
		"Period checked:\n"+
		"From:\t%s\n"+
		"To:\t%s\n", d.Metrics[0].Date, d.Metrics[len(d.Metrics)-1].Date)
}
