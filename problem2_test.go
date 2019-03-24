package main

import (
	"reflect"
	"testing"
	"time"
)

func Test_problem2(t *testing.T) {
	type args struct {
		taskCount int
	}
	tests := []struct {
		name string
		args args
		want []time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problem2(tt.args.taskCount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("problem2() = %v, want %v", got, tt.want)
			}
		})
	}
}
