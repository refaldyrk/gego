package helper

import "strings"

func Include(s []string, sub string) bool {
	for _, v := range s {
		if strings.Contains(v, sub) {
			return true
		}
	}
	return false
}

func Join(s []string, sep string) string {
	return strings.Join(s, sep)
}

func Split(s string, sep string) []string {
	return strings.Split(s, sep)
}
