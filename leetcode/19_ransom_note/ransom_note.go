package ransom_note

func canConstruct(ransomNote string, magazine string) bool {
	charCount := make(map[rune]int)
	for _, char := range magazine {
		charCount[char]++
	}
	for _, char := range ransomNote {
		charCount[char]--
		if charCount[char] < 0 {
			return false
		}
	}
	return true
}
