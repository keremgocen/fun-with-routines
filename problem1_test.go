package main

import (
	"testing"
)

func Test_problem1(t *testing.T) {
	type args struct {
		taskCount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"correct number of random values are generated",
			args{
				taskCount: 100,
			},
			100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := problem1(tt.args.taskCount)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

