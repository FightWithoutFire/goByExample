package main

import (
	"fmt"
	"testing"
)

func TestIntMinBasic(t *testing.T) {
	min := IntMin(2, -2)

	if min != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2", min)
	}
}
func TestMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{{1, 2, 1}, {2, 3, 2}}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}

}
func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}
