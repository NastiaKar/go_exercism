package isogram

import "unicode"

func IsIsogram(word string) bool {
	seen := make(map[rune]bool)
	for _, ch := range word {
        if !unicode.IsLetter(ch) {
            continue
        }
        ch = unicode.ToLower(ch)
        if seen[ch] {
            return false
        }
        seen[ch] = true
    }
    return true
}
