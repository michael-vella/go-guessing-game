package main

import (
	"fmt"
	"testing"
)

func TestGetLives(t *testing.T) {
	var tests = []struct {
		in   int
		want int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 2},
		{5, 3},
		{10, 4},
		{100, 7},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d", tt.in)
		t.Run(testName, func(t *testing.T) {
			ans := GetLives(tt.in)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
