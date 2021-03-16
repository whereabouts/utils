package code

import (
	"fmt"
	"math/rand"
	"strings"
)

func String() string {
	return StrGenerate(DefaultLen, TypeBlend)
}

func NewStr(length int, codeType int) string {
	return StrGenerate(length, codeType)
}

func StrBlend(length int) string {
	return StrGenerate(length, TypeBlend)
}

func StrDigit(length int) string {
	return StrGenerate(length, TypeDigit)
}

func StrLetterLower(length int) string {
	return StrGenerate(length, TypeLetterLower)
}

func StrLetterUpper(length int) string {
	return StrGenerate(length, TypeLetterUpper)
}

func StrGenerate(length int, codeType int) (code string) {
	for i := 0; i < length; i++ {
		switch codeType {
		case TypeBlend:
			code = fmt.Sprintf("%s%c", code, characterBlend[rand.Intn(len(characterBlend))])
		case TypeDigit:
			code = fmt.Sprintf("%s%c", code, characterDigit[rand.Intn(len(characterDigit))])
		case TypeLetterLower:
			code = fmt.Sprintf("%s%c", code, characterLetterLower[rand.Intn(len(characterLetterLower))])
		case TypeLetterUpper:
			code = fmt.Sprintf("%s%c", code, characterLetterUpper[rand.Intn(len(characterLetterUpper))])
		default:
			code = fmt.Sprintf("%s%c", code, characterBlend[rand.Intn(len(characterBlend))])
		}
	}
	return code
}

func VerifyStr(in string, code string, capslock ...bool) bool {
	if len(capslock) == 0 || !capslock[0] {
		return strings.ToUpper(in) == strings.ToUpper(code)
	} else {
		return in == code
	}
}
