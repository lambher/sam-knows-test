
<div align="center">
  <img src="https://samknows.com/img/sk-logo.svg" align="center" width="60">
  <h1 align="center">SamKnows Backend Engineering Test</h1>
</div>

## A brief description of the project

The scope of the test is to generate the expected output files (in the `outputs` folder) given the input files (in the `inputs` folder).

The application should:
1. Display the min, max, median and average for a data set.
2. Discover under-performing periods of download performance.

## Installation instructions

Before running the test, please make sure you have Golang installed on your OS.
https://golang.org/doc/install

## Example

Let says you have an input file `inputs/1.json` that contains:

```[
  {
    "metricValue": 12693166.98,
    "dtime": "2018-01-29"
  },
  {
    "metricValue": 12668239.57,
    "dtime": "2018-01-30"
  },
  {
    "metricValue": 12723772.1,
    "dtime": "2018-01-31"
  },
...
]
```

Just run the following command : `go run main.go inputs/1.json`

You should have this output :

```
SamKnows Metric Analyser v1.0.0
===============================

Period checked:

    From: 2018-01-29
    To:   2018-02-27

Statistics:

    Unit: Megabits per second

    Average: 102.7
    Min: 101.25
    Max: 104.08
    Median: 102.93
```

## Run the tests

To run the tests and show the percent of coverage statements please run : `go test ./... -cover`
