package processor

import "strings"

func applyArticleRule(tokens []string) []string {
	vowels := "aeiouhAEIOUH"
	for i := 0; i < len(tokens)-1; i++ {
		if tokens[i] == "a" && strings.ContainsRune(vowels, rune(tokens[i+1][0])) {
			tokens[i] = "an"
	}
		if tokens[i] == "A" && strings.ContainsRune(vowels, rune(tokens[i+1][0])) {
			tokens[i] = "An"
	}
	}
	return tokens
}