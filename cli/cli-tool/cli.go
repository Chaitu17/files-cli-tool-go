package cli

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func FileReaderCLIHandler(bytes *bool, lines *bool, words *bool, chars *bool, fileContent string) {
	switch {
	case *bytes:
		fmt.Println("bytes", CountNoOfBytes(fileContent))
	case *lines:
		fmt.Println("lines", CountNoOfLines(fileContent))
	case *words:
		fmt.Println("words", CountNoOfWords(fileContent))
	case *chars:
		fmt.Println("characters", CountNoOfCharacters(fileContent))
	default:
		fmt.Println("lines", CountNoOfLines(fileContent))
		fmt.Println("words", CountNoOfWords(fileContent))
		fmt.Println("bytes", CountNoOfBytes(fileContent))
	}
}

func CountNoOfBytes(fileContent string) int {
	return len(fileContent)
}

func CountNoOfLines(fileContent string) int {
	return strings.Count(fileContent, "\r\n")
}

func CountNoOfWords(fileContent string) int {
	return len(strings.Fields(fileContent))
}

func CountNoOfCharacters(fileContent string) int {
	return utf8.RuneCountInString(fileContent)
}
