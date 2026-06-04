package processor

import "strings"
import "unicode"

func applyCaseRules(tokens []string) []string {
	for i := 0; i < len(tokens); i++ {
		op, count := parseCaseToken(tokens[i])
		if op == "" {
			continue
	}
		
		for j := 1; j <= count && i-j >= 0; j++ {
			idx := i - j
			switch op {
			case "up":
				tokens[idx] = strings.ToUpper(tokens[idx])
			case "low":
				tokens[idx] = strings.ToLower(tokens[idx])
			case "cap":
				tokens[idx] = capitalize(tokens[idx])
			}
	}
		tokens = remove(tokens, i)
		i--
	}
	return tokens
}

func capitalize(s string) string {
	if s == "" {
		return s
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	for i := 1; i < len(r); i++ {
		r[i] = unicode.ToLower(r[i])
	}
	return string(r)
}

func parseCaseToken(tok string) (string, int) {
	if tok == "(up)" { return "up", 1 }
	if tok == "(low)" { return "low", 1 }
	if tok == "(cap)" { return "cap", 1 }
	
	if strings.HasPrefix(tok, "(up,") { return "up", parseNum(tok) }
	if strings.HasPrefix(tok, "(low,") { return "low", parseNum(tok) }
	if strings.HasPrefix(tok, "(cap,") { return "cap", parseNum(tok) }
	return "", 0
}