package main

import "testing"

func Test_readFile(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "test1",args: args{"puzzle.txt"},want:1598415 },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readFile(tt.args.s); got != tt.want {
				t.Errorf("readFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
