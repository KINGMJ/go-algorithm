package validpalindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
		{".,", true},
		{"0P", false},
	}

	for _, tc := range testCases {
		result := IsPalindrome(tc.input)
		if result != tc.expected {
			t.Errorf("IsPalindrome(%q) = %v, 期望 %v", tc.input, result, tc.expected)
		}
	}
}
