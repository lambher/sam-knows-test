package models

import (
	"testing"
)

func TestBytesToMbits(t *testing.T) {
	testValues := []struct {
		Value    float64
		Expected float64
	}{
		{
			Value:    4583,
			Expected: 0.036664,
		},
		{
			Value:    9498464654,
			Expected: 75987.717232,
		},
		{
			Value:    98465465431,
			Expected: 787723.723448,
		},
		{
			Value:    0,
			Expected: 0,
		},
	}

	for _, testValue := range testValues {
		result := bytesToMbits(testValue.Value)
		if result != testValue.Expected {
			t.Errorf("Conversion of %f bytes to Mbits was incorrect, got: %f, want: %f.", testValue.Value, result, testValue.Expected)
		}
	}
}

func TestFindMedian(t *testing.T) {
	testValues := []struct {
		Values   []float64
		Expected float64
	}{
		{
			Values:   []float64{1, 9, 2.3, 10, -5, 4},
			Expected: 3.15,
		},
		{
			Values:   []float64{1, 2, 2.5, 6, 10, 24, 45},
			Expected: 6,
		},
		{
			Values:   []float64{},
			Expected: 0,
		},
		{
			Values:   []float64{9},
			Expected: 9,
		},
		{
			Values:   []float64{42, 42},
			Expected: 42,
		},
		{
			Values:   []float64{42, 0},
			Expected: 21,
		},
	}

	for _, testValue := range testValues {
		result := findMedian(testValue.Values)
		if result != testValue.Expected {
			t.Errorf("FindMedian result was incorrect, got: %f, want: %f.", result, testValue.Expected)
		}
	}
}

func TestLoadDataFromFile(t *testing.T) {
	testValues := []struct {
		Path        string
		LenExpected int
	}{
		{
			Path:        "../inputs/1.json",
			LenExpected: 30,
		},
		{
			Path:        "../inputs/2.json",
			LenExpected: 30,
		},
		{
			Path:        "../inputs/3.json",
			LenExpected: 15,
		},
	}

	for _, testValue := range testValues {
		data, err := LoadDataFromFile(testValue.Path)
		if err != nil {
			t.Error(err)
			continue
		}
		if len(data.Metrics) != testValue.LenExpected {
			t.Errorf("metrics len was incorrect, got: %d, want: %d.", len(data.Metrics), testValue.LenExpected)
		}
	}
}

func TestData_Process(t *testing.T) {
	testValues := []struct {
		Path         string
		DataExpected Data
	}{
		{
			Path: "../inputs/1.json",
			DataExpected: Data{
				Average: 1.2837656825666666e+07,
				Min:     1.265595155e+07,
				Max:     1.300967374e+07,
				Median:  1.286665518e+07,
			},
		},
		{
			Path: "../inputs/2.json",
			DataExpected: Data{
				Average: 1.19370179814e+07,
				Min:     3.453456324e+06,
				Max:     1.300967374e+07,
				Median:  1.2863517795e+07,
			},
		},
		{
			Path: "../inputs/3.json",
			DataExpected: Data{
				Average: 10692567.2,
				Min:     1781257.37,
				Max:     12991749.01,
				Median:  12907999.21,
			},
		},
	}

	for _, testValue := range testValues {
		data, err := LoadDataFromFile(testValue.Path)
		if err != nil {
			t.Error(err)
			continue
		}
		err = data.Process()
		if err != nil {
			t.Error(err)
			continue
		}
		if data.ResultToString() != testValue.DataExpected.ResultToString() {
			t.Errorf("data was incorrect, got: \n%s, want: \n%s.", data.ResultToString(), testValue.DataExpected.ResultToString())
		}
	}
}

func TestFindUnderPerformingPeriods(t *testing.T) {
	testValues := []struct {
		Path         string
		DataExpected Data
	}{
		{
			Path: "../inputs/1.json",
			DataExpected: Data{
				UnderPerformingPeriods: []Period{},
			},
		},
		{
			Path: "../inputs/2.json",
			DataExpected: Data{
				UnderPerformingPeriods: []Period{
					{
						Start: NewDTime("2018-02-05"),
						End:   NewDTime("2018-02-07"),
					},
				},
			},
		},
		{
			Path: "../inputs/3.json",
			DataExpected: Data{
				UnderPerformingPeriods: []Period{
					{
						Start: NewDTime("2018-02-25"),
						End:   NewDTime("2018-02-27"),
					},
				},
			},
		},
	}

	for _, testValue := range testValues {
		data, err := LoadDataFromFile(testValue.Path)
		if err != nil {
			t.Error(err)
			continue
		}
		err = data.Process()
		if err != nil {
			t.Error(err)
			continue
		}
		if len(data.UnderPerformingPeriods) != len(testValue.DataExpected.UnderPerformingPeriods) {
			t.Errorf("UnderPerformingPeriods's len was incorrect, got: \n%d, want: \n%d.", len(data.UnderPerformingPeriods), len(testValue.DataExpected.UnderPerformingPeriods))
			continue
		}
		for i, underPerformingPeriod := range data.UnderPerformingPeriods {
			if !underPerformingPeriod.Start.ToTime().Equal(testValue.DataExpected.UnderPerformingPeriods[i].Start.ToTime()) {
				t.Errorf("UnderPerformingPeriod's start was incorrect, got: \n%s, want: \n%s.", underPerformingPeriod.Start, testValue.DataExpected.UnderPerformingPeriods[i].Start)
			}
			if !underPerformingPeriod.End.ToTime().Equal(testValue.DataExpected.UnderPerformingPeriods[i].End.ToTime()) {
				t.Errorf("UnderPerformingPeriod's end was incorrect, got: \n%s, want: \n%s.", underPerformingPeriod.End, testValue.DataExpected.UnderPerformingPeriods[i].End)
			}
		}
	}
}
