package gego

import "github.com/refaldyrk/gego/helper"

func Result(s string) string {
	word := helper.Split(s, "")
	vocal := []string{"a", "i", "u", "e", "o"}
	var last []string

	for i := 0; i < len(word); i++ {
		if helper.Include(vocal, word[i]) {
			last = append(last, word[i]+"g"+word[i])
		} else {
			last = append(last, word[i])
		}
	}

	return helper.Join(last, "")
}
