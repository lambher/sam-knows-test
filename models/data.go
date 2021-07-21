package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
)

type Data struct {
	Metrics []Metric

	Average float64
	Min     float64
	Max     float64
	Median  float64

	UnderPerformingPeriods []Period
}

type Period struct {
	Start DTime
	End   DTime
}

type Metric struct {
	Value float64 `json:"metricValue"`
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

func (d *Data) Process() error {
	if len(d.Metrics) == 0 {
		return errors.New("no metrics to process")
	}

	d.Average = 0
	d.Min = d.Metrics[0].Value
	d.Max = d.Metrics[0].Value

	values := make([]float64, 0)

	for _, metric := range d.Metrics {
		d.Average += metric.Value
		if d.Min > metric.Value {
			d.Min = metric.Value
		}
		if d.Max < metric.Value {
			d.Max = metric.Value
		}
		values = append(values, metric.Value)
	}
	d.Average /= float64(len(d.Metrics))
	d.Median = findMedian(values)

	d.findUnderPerformingPeriods()
	return nil
}

func (d *Data) findUnderPerformingPeriods() {
	gap := 5. * 125000

	if d.Min > d.Median-gap {
		return
	}

	underPerforming := false
	for _, metric := range d.Metrics {
		if metric.Value < d.Median-gap {
			if !underPerforming {
				d.UnderPerformingPeriods = append(d.UnderPerformingPeriods, Period{
					Start: metric.Date,
				})
			}
			d.UnderPerformingPeriods[len(d.UnderPerformingPeriods)-1].End = metric.Date
			underPerforming = true
		} else {
			underPerforming = false
		}

	}
}

func findMedian(values []float64) float64 {
	sort.Float64s(values)
	if len(values)%2 == 1 {
		return values[len(values)/2]
	}
	return (values[len(values)/2] + values[len(values)/2-1]) / 2
}

func (d *Data) String() string {
	if len(d.Metrics) == 0 {
		return "" +
			"SamKnows Metric Analyser v1.0.0\n" +
			"===============================\n" +
			"\n" +
			"No metrics processed"
	}

	str := fmt.Sprintf(""+
		"SamKnows Metric Analyser v1.0.0\n"+
		"===============================\n"+
		"\n"+
		"Period checked:\n"+
		""+
		"\tFrom:\t%s\n"+
		"\tTo:\t%s\n"+
		"\n"+
		"Statistics:\n"+
		"\n"+
		"\tUnit: Megabits per second\n"+
		"\n"+
		"\tAverage: %.1f\n"+
		"\tMin: %.2f\n"+
		"\tMax: %.2f\n"+
		"\tMedian: %.2f\n",
		d.Metrics[0].Date,
		d.Metrics[len(d.Metrics)-1].Date,
		bytesToMbits(d.Average),
		bytesToMbits(d.Min),
		bytesToMbits(d.Max),
		bytesToMbits(d.Median))

	if len(d.UnderPerformingPeriods) > 0 {
		str += fmt.Sprintf("\n" +
			"Under-performing periods:\n")
	}

	for _, underPerformingPeriod := range d.UnderPerformingPeriods {
		str += fmt.Sprintf("\n"+
			"\t* The period between %s and %s\n"+
			"\t  was under-performing.\n", underPerformingPeriod.Start, underPerformingPeriod.End)
	}

	return str
}

//1 Mbit = 125000 bytes
func bytesToMbits(value float64) float64 {
	return value / 125000
}
