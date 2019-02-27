package main

import "testing"

func TestMinMax(t *testing.T) {
    tests := []struct {
        numbers  []float64
        min, max float64
    }{
        {
            numbers: []float64{1, 2, 3, 4, 5},
            min:     1,
            max:     5,
        },
        {
            numbers: []float64{},
            min:     0,
            max:     0,
        },
    }

    for _, test := range tests {
        min := min(test.numbers)
        max := max(test.numbers)

        if min != test.min {
            t.Errorf("Expected: %f - Result: %f", test.min, min)
        }

        if max != test.max {
            t.Errorf("Expected: %f - Result: %f", test.max, max)
        }
    }
}
