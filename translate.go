package piglatin

import (
	"strings"
)

const (
	pigLatinSuffix             string = "ay"
	firstLetterExceptions      string = "aeiou"
	firstLetterExceptionSuffix string = "d" + pigLatinSuffix
)

// Translate translates an English word into Pig latin.
func Translate(in string) string {
	var latinWords []string
	englishWords := strings.Split(in, " ")

	for _, word := range englishWords {
		first := word[0:1]
		var latinWord string
		if strings.Contains(firstLetterExceptions, first) {
			latinWord = word + firstLetterExceptionSuffix
		} else {
			latinWord = word[1:] + first + pigLatinSuffix
		}
		latinWords = append(latinWords, latinWord)
	}
	return strings.Join(latinWords, " ")
}
