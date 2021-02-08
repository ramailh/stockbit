package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "a(test eko) asdfasdf"
	fmt.Println(newFindStringInBracket(str))

}

// New
func newFindStringInBracket(str string) string {
	if len(str) > 0 {
		indexOpenBracket := strings.Index(str, "(")
		indexCloseBracket := strings.Index(str, ")")

		var arrayStr []string
		for i, letter := range str {
			if i > indexOpenBracket && i < indexCloseBracket {
				arrayStr = append(arrayStr, string(letter))
			}
		}

		str = strings.Join(arrayStr, "")
	}

	return str
}

// OLD
func findFirstStringInBracket(str string) string {
	if len(str) > 0 {
		indexFirstBracketFound := strings.Index(str, "(")
		fmt.Println(indexFirstBracketFound)
		if indexFirstBracketFound >= 0 {
			runes := []rune(str)
			wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
			indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
			if indexClosingBracketFound >= 0 {
				runes := []rune(wordsAfterFirstBracket)
				return string(runes[1 : indexClosingBracketFound-1])
			} else {
				return ""
			}
		} else {
			return ""
		}
	} else {
		return ""
	}
	return ""
}
