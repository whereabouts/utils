package code

import (
	"fmt"
	"math/rand"
	"strings"
)

func Custom(template string) string {
	return NewCustom(DefaultLen, template)
}

func NewCustom(length int, template string) (code string) {
	for i := 0; i < length; i++ {
		index := rand.Intn(len(template) / 3)
		for j, v := range template {
			if j == index*3 {
				code = fmt.Sprintf("%s%c", code, v)
				break
			}
		}
	}
	return code
}

func VerifyCustom(in string, code string, capslock ...bool) bool {
	if len(capslock) == 0 || !capslock[0] {
		return strings.ToUpper(in) == strings.ToUpper(code)
	} else {
		return in == code
	}
}
