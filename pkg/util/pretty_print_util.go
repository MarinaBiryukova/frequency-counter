package util

import (
	"fmt"
	"strconv"
	"unicode"
)

func PrettyPrint(key string, value int) {
	unicodeKey := []rune(key)
	fmt.Println(string(unicode.ToUpper(unicodeKey[0])) + string(unicodeKey[1:]) + ": " + strconv.Itoa(value))
}
