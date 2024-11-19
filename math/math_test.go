package math

import "testing"

type averagetestpair struct {
	values  []float64
	average float64
}

type mintestpair struct {
	values []float64
	min    float64
}

type maxtestpair struct {
	values []float64
	max    float64
}

var averagetests = []averagetestpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}

var mintest = []mintestpair{
	{[]float64{1, 2, 3}, 1},
	{[]float64{6, 5, 3}, 3},
	{[]float64{10, 7, 9}, 7},
}

var maxtest = []maxtestpair{
	{[]float64{1, 2, 3}, 3},
	{[]float64{6, 5, 3}, 6},
	{[]float64{10, 7, 9}, 10},
}

func TestAverage(t *testing.T) {
	for _, pair := range averagetests {
		v := Average(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}

func TestMin(t *testing.T) {
	for _, pair := range mintest {
		v := Min(pair.values...)
		if v != pair.min {
			t.Error(
				"For", pair.values,
				"expected", pair.min,
				"got", v,
			)
		}
	}
}

func TestMax(t *testing.T) {
	for _, pair := range maxtest {
		v := Max(pair.values...)
		if v != pair.max {
			t.Error(
				"For", pair.values,
				"expected", pair.max,
				"got", v,
			)
		}
	}
}
