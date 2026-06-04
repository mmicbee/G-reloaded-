package processor

import "strings"

// fixPunctuation handles: spacing before,.!?:;, groups like... and!?, and quotes ' word '
func fixPunctuation(text string) string {
	// 1. Fix quotes first: ' word ' -> 'word'
	text = fixQuotes(text)
	
	// 2. Fix spacing around punctuation
	// Remove space before punctuation
	punct := []string{", ".", "!", "?", ":", ";"}
	for _, p := range punct {
		text = strings.ReplaceAll(text, " "+p, p)
	}
	
	// Ensure space after punctuation, except for groups... and!? etc
	text = addSpaceAfterPunct(text)
	
	return text
}

// fixQuotes: handles ' word ' and ' multi word phrase '
func fixQuotes(text string) string {
	var result strings.Builder
	inQuote := false
	
	words := strings.Fields(text) // split by any whitespace
	
	for i := 0; i < len(words); i++ {
		w := words[i]
		
		if w == "'" {
			if!inQuote {
				// opening quote
				result.WriteString("'")
				inQuote = true
			} else {
				// closing quote
				result.WriteString("'")
				inQuote = false
				// add space after closing quote if not end
				if i < len(words)-1 {
					result.WriteString(" ")
				}
			}
			continue
	}
		
		if inQuote {
			result.WriteString(w)
			// add space between words inside quotes
			if i < len(words)-1 && words[i+1]!= "'" {
				result.WriteString(" ")
			}
	} else {
			result.WriteString(w)
			if i < len(words)-1 {
				result.WriteString(" ")
			}
	}
	}
	return result.String()
}

// addSpaceAfterPunct: add space after.,!? : ; but skip groups like... or!?
func addSpaceAfterPunct(text string) string {
	var result strings.Builder
	
	for i := 0; i < len(text); i++ {
		c := text[i]
		result.WriteByte(c)
		
		if c == '.' || c == ',' || c == '!' || c == '?' || c == ':' || c == ';' {
			// Check if next char exists and is not space and not same punct = group
			if i+1 < len(text) {
				next := text[i+1]
				if next!= ' && next!= c {
					// single punctuation, add space
					result.WriteByte(' ')
				}
			}
	}
	}
	return result.String()
}