package utils

import (
	"strings"
)

func StrRepeat(text string, midText string, n int) string {
	return strings.Repeat("\n", 1) + strings.Repeat(text, n) + strings.Repeat(midText, 1) + strings.Repeat(text, n) + strings.Repeat("\n\n", 1)
}
