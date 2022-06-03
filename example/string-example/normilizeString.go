package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

func main() {

	str := "companies&markets news"
	//str := "companies"
	res := normalizeSectionName(str)
	fmt.Println("result : ", res)
}

func normalizeSectionName(feed string) string {
	var result = strings.TrimSpace(feed)
	result = strings.ToLower(result)

	result = strings.ReplaceAll(result, "&", " ")
	result = strings.ReplaceAll(result, ",", " ")
	result = strings.ReplaceAll(result, "-", " ")

	// remove duplicate whitespace from a string
	space := regexp.MustCompile(`\s+`)
	result = space.ReplaceAllString(result, " ")

	result = strings.ReplaceAll(result, " ", "_")

	return result
}

func hasSymbol(str string) bool {
	for _, letter := range str {
		if unicode.IsSymbol(letter) {
			return true
		}
	}
	return false
}
