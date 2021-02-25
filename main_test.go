package main

import "testing"

func Add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"1+2", 1, 2, 3},
		{"2+5", 2, 5, 7},
		{"99+1", 99, 1, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if tt.want != got {
				t.Errorf("unexpected result. want%d, got=%d", tt.want, got)
			}
		})
	}
}
