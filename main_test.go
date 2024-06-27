package main

import "testing"

func TestFoo(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			"Hello World Test",
			"Hello, World!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := foo(); got != tt.want {
				t.Errorf("foo() = %v, want %v", got, tt.want)
			}
		})
	}
}
