package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	banner("Go", 6)
	banner("G☺", 6)

	s := "G☺"
	fmt.Println("len:", len(s))
	// code point = rune ~= unicode character
	for i, r := range s {
		fmt.Println(i, r)
		if i == 0 {
			fmt.Printf("%c of type %T\n", r, r)
			// rune (int32)
		}
	}

	b := s[0]
	fmt.Printf("%c of type %T\n", b, b)
	// byte (uint8)
	x, y := 1, "1"
	fmt.Printf("x=%v, y=%v\n", x, y)
	fmt.Printf("x=%#v, y=%#v\n", x, y) // Use #v in debug/log
	fmt.Printf("%20s!\n", s)
	fmt.Println("g", isPalindrome("g"))
	fmt.Println("go", isPalindrome("go"))
	fmt.Println("gog", isPalindrome("gog"))
	fmt.Println("g☺g", isPalindrome("g☺g"))
	fmt.Println("gogo", isPalindrome("gogo"))
}

// isPalindrome("g") -> true
// isPalindrome("go") -> false
// isPalindrome("gog") -> true
// isPalindrome("gogo") -> false
func isPalindrome(s string) bool {
	runeStr := []rune(s) // get slice of rune out of s
	var reversedRuneStr []rune
	for i := (len(runeStr) - 1); i >= 0; i-- { // create the slice in reverse order
		reversedRuneStr = append(reversedRuneStr, runeStr[i])
	}
	return string(runeStr) == string(reversedRuneStr) //compare both slices converted into string
}

func banner(text string, width int) {
	padding := (width - utf8.RuneCountInString(text))
	// padding := (width - len(text)) / 2 BUG: len is in bytes
	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}
	fmt.Println(text)
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}
