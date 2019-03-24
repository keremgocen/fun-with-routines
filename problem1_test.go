package main

import (
	"testing"
)

func Test_problem1(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			"correct number of random values generated",
			100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := problem1()
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

