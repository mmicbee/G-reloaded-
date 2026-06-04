package processor

import (
	"strconv"
	"strings" // <- add this
)

func remove(slice []string, i int) []string {
	if i < 0 || i >= len(slice) {
		return slice
	}
	return append(slice[:i], slice[i+1:]...)
}

func parseNum(tok string) int {
	// "(up, 3)" -> 3
	start := strings.Index(tok, ",") + 2
	end := len(tok) - 1
	if start >= end {
		return 1
	}
	n, _ := strconv.Atoi(tok[start:end])
	return n
}