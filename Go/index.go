package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	str1 := "kUpu-Kupu"
	fmt.Println("Is", "\"", str1, "\"", "palindrome?", palindrome(str1))
	str2 := "kAsur ini Rusak"
	fmt.Println("Is", "\"", str2, "\"", "palindrome?", palindrome(str2))
}

func palindrome(str string) bool {
	re := regexp.MustCompile("[^a-zA-Z0-9]")
	str = re.ReplaceAllLiteralString(str, "")
	str = strings.ToLower(str)

	// Kamu hanya perlu mengecek setengah dari isi stringnya saja
	// tidak perlu semuanya
	len := len(str)
	for i := 0; i < len/2; i++ {
		if str[i] != str[len-i-1] {
			return false
		}
	}

	return true
}
