package pangram

import "strings"

func IsPangram(input string) bool {
	seen := make(map[rune]bool)
    for _, ch := range strings.ToLower(input) {
        if ch >= 'a' && ch <= 'z' {
            seen[ch] = true
        }
    }
    return len(seen) == 26
}
