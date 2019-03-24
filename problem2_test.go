package main

import (
	"testing"
)

func Test_problem2(t *testing.T) {
	type args struct {
		taskCount int
	}
	tests := []struct {
		name string
		args args
		low  float64
		high float64
	}{
		{
			"difference between timestamps are between 0.9 and 1.1 seconds",
			args{
				taskCount: 5,
			},
			0.9,
			1.1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := problem2(tt.args.taskCount)
			for i, v := range ts {
				if i+1 < tt.args.taskCount {
					delta := ts[i+1].Sub(v)
					got := delta.Seconds()
					if got < tt.low || got > tt.high {
						t.Errorf("problem2(%d) got:%v should be between %v < x < %v", i, got, tt.low, tt.high)
					}
				}
			}
		})
	}
}
