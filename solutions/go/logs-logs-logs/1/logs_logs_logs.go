package logs

import (
    "unicode/utf8";
    "strings"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	var applicationFormat = map[rune]string{
        '❗' : "recommendation",
        '🔍' : "search",
        '☀' : "weather",
    }

    for _, c := range log {
        val, ok := applicationFormat[c]
        if ok {
            return val
        }
    }
    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var sb strings.Builder
    for _, c := range log {
        if c == oldRune {
            sb.WriteRune(newRune)
        } else {
            sb.WriteRune(c)
        }
    }

    return sb.String()
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
