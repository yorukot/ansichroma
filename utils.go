package ansichroma

import (
	"strings"

	"github.com/alecthomas/chroma/v2"
)

func getTrileanToBool(input chroma.Trilean) bool {
	switch input {
	case chroma.Yes:
		return true
	case chroma.No:
		return false
	default:
		return false
	}
}

func trimTrailingNewlines(s string) (string, int) {
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '\n' {
			count++
		} else {
			break
		}
	}
	trimmed := strings.TrimRight(s, "\n")
	return trimmed, count
}

func truncateString(s string, maxLength int) string {
    if len(s) > maxLength {
        return s[:maxLength]
    }
    return s
}