package main

import "testing"

// TestIsEven verifies that isEven correctly identifies even and odd integers,
// including zero, positive, negative, and large numbers.
func TestIsEven(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{0, true},   // zero is even
		{1, false},  // positive odd
		{2, true},   // positive even
		{-3, false}, // negative odd
		{-4, true},  // negative even
		{100, true}, // large even number
	}
	for _, tt := range tests {
		if got := isEven(tt.n); got != tt.want {
			t.Errorf("isEven(%d) = %v, want %v", tt.n, got, tt.want)
		}
	}
}

// TestAdd verifies that add correctly sums two integers across various cases
// including positives, zeros, negatives, and mixed signs.
func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{10, 20, 30},    // basic positive sum
		{0, 0, 0},       // both zeros
		{-5, 5, 0},      // cancels to zero
		{-3, -7, -10},   // both negative
		{100, 200, 300}, // large numbers
	}
	for _, tt := range tests {
		if got := add(tt.a, tt.b); got != tt.want {
			t.Errorf("add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

// TestCheckString verifies that checkString correctly classifies strings
// as empty, numeric, alphanumeric, or neither.
func TestCheckString(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"", "empty"},              // empty string
		{"12345", "numeric"},       // digits only
		{"abc123", "alphanumeric"}, // letters and digits mixed
		{"hello", "alphanumeric"},  // letters only
		{"!@#$", "neither"},        // special characters only
		{"abc!@#", "neither"},      // letters mixed with special chars
		{"0", "numeric"},           // single digit
		{"ABC", "alphanumeric"},    // uppercase letters
	}
	for _, tt := range tests {
		if got := checkString(tt.s); got != tt.want {
			t.Errorf("checkString(%q) = %q, want %q", tt.s, got, tt.want)
		}
	}
}
