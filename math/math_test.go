package math

import "testing"

type testpair[T any] struct {
	values []float64
	result T
}

var averagetests = []testpair[float64]{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}

var mintests = []testpair[float64]{
	{[]float64{1, 2, 3}, 1},
	{[]float64{6, 5, 3}, 3},
	{[]float64{10, 7, 9}, 7},
}

var maxtests = []testpair[float64]{
	{[]float64{1, 2, 3}, 3},
	{[]float64{6, 5, 3}, 6},
	{[]float64{10, 7, 9}, 10},
}

func TestAverage(t *testing.T) {
	runTests(t, averagetests, Average)
}

func TestMin(t *testing.T) {
	runTests(t, mintests, func(values []float64) float64 {
		return Min(values...)
	})
}

func TestMax(t *testing.T) {
	runTests(t, maxtests, func(values []float64) float64 {
		return Max(values...)
	})
}

func runTests[T comparable](t *testing.T, tests []testpair[T], f func([]float64) T) {
	for _, pair := range tests {
		if v := f(pair.values); v != pair.result {
			t.Errorf("For %v, expected %v, got %v", pair.values, pair.result, v)
		}
	}
}
