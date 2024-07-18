package main

import "testing"

func TestCountVowels(t *testing.T) {
	testCases := []struct {
		userstring    string
		expected int
	}{
		{"vowelcount", 4},
		{"aeiou", 5},
		{"AEIOU", 5},
		{"", 0},
		{"#a@e$i", 3},
		{"Hello, How are you", 7},
	}

	for _, tc := range testCases {
		result := CountVowels(tc.userstring)
		if result != tc.expected {
			t.Errorf("CountVowels(%q) = %d; expected %d", tc.userstring, result, tc.expected)
		}
	}
}
