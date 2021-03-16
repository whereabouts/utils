package code

import (
	"fmt"
	"math/rand"
)

func PrintCNChar() {
	for _, v := range zn_cn {
		fmt.Printf("%c", v)
	}
}

func ZN_CN() string {
	return NewZN_CN(DefaultLen)
}

func NewZN_CN(length int) (code string) {
	for i := 0; i < length; i++ {
		index := rand.Intn(len(zn_cn) / 3)
		for j, v := range zn_cn {
			if j == index*3 {
				code = fmt.Sprintf("%s%c", code, v)
				break
			}
		}
	}
	return code
}

func VerifyZN_CN(in string, code string) bool {
	return in == code
}
