package main

import "testing"

func TestIsEven(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{0, true},
		{1, false},
		{2, true},
		{-3, false},
		{-4, true},
		{100, true},
	}
	for _, tt := range tests {
		if got := isEven(tt.n); got != tt.want {
			t.Errorf("isEven(%d) = %v, want %v", tt.n, got, tt.want)
		}
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{10, 20, 30},
		{0, 0, 0},
		{-5, 5, 0},
		{-3, -7, -10},
		{100, 200, 300},
	}
	for _, tt := range tests {
		if got := add(tt.a, tt.b); got != tt.want {
			t.Errorf("add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestCheckString(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"", "empty"},
		{"12345", "numeric"},
		{"abc123", "alphanumeric"},
		{"hello", "alphanumeric"},
		{"!@#$", "neither"},
		{"abc!@#", "neither"},
		{"0", "numeric"},
		{"ABC", "alphanumeric"},
	}
	for _, tt := range tests {
		if got := checkString(tt.s); got != tt.want {
			t.Errorf("checkString(%q) = %q, want %q", tt.s, got, tt.want)
		}
	}
}
