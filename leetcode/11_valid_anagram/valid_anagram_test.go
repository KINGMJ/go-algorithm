package validanagram

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"", "", true},
		{"a", "a", true},
		{"ab", "ba", true},
		{"ab", "ab", true},
		{"ab", "abc", false},
		{"abc", "ab", false},
		{"aacc", "ccac", false},
		{"汉字", "字汉", true},
	}

	for _, test := range tests {
		result := isAnagram(test.s, test.t)
		if result != test.expected {
			t.Errorf("isAnagram(%q, %q) = %v; want %v", test.s, test.t, result, test.expected)
		}
	}
}
