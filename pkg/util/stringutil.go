package util

import "strings"

func SplitRunes(s string,runes []rune) [] string {
	f := strings.FieldsFunc(s, func(r rune) bool {
		for _, rr := range runes {
			if rr == r {
				return true
			}
		}
		return false
	})
	return f
}
