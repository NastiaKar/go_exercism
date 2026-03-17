package luhn

import "strings"

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")

	if len(id) <= 1 {
		return false
	}

	var sum int
	for i, ch := range id {
		if ch < '0' || ch > '9' {
			return false
		}
		num := int(ch - '0')
		if (len(id)-1-i)%2 == 1 {
			num *= 2
			if num > 9 {
				num -= 9
			}
		}
		sum += num
	}

	return sum%10 == 0
}