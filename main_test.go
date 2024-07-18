package main

import "testing"

func TestCountVowels(t *testing.T) {
	testCases := []struct {
		userstring    string
		expected int
	}{
        {"vowelcount", 4},         // Test case with mixed letters
        {"aeiou", 5},              // Test case with all vowels
        {"AEIOU", 5},              // Test case with uppercase vowels
        {"", 0},                   // Test case with an empty string
        {"#a@e$i", 3},             // Test case with special characters and vowels
        {"Hello, How are you", 7}, // Test case with a sentence
	}

	for _, tc := range testCases {
		result := CountVowels(tc.userstring)
		if result != tc.expected {
			t.Errorf("CountVowels(%q) = %d; expected %d", tc.userstring, result, tc.expected)
		}
	}
}
