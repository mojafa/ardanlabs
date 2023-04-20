package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	banner("Go", 6)

	banner("G😝", 6)

	s := "G😝"
	fmt.Println("len :", len(s))
	//code point = rune ~=unicode

	for i, r := range s {
		fmt.Println(i, r)
		if i == 0 {
			fmt.Printf("%c of type %T\n", r, r)
			//rune =int32

		}
	}

	b := s[0]
	fmt.Printf("%c of type %T\n", b, b)
	//byte =uint8

	x, y := 1, "1"
	fmt.Printf("x=%v, y=%v\n", x, y)

	//use #v in debugging or log
	fmt.Printf("x=%#v, y=%#v\n", x, y)

	fmt.Printf("%20s!\n", s)

	fmt.Println("go", isPalindrome("go"))

}

// isPalindrome "g" --> true
// isPalindrome "go" --> false
// isPalindrome "gog" --> true
func isPalindrome(s string) bool {
	//to make it unicode aware

	rs := []rune(s)
	for i := 0; i < len(rs)/2; i++ {
		if rs[i] != rs[len(rs)-i-1] {
			return false
		}
	}
	return true
}
func banner(text string, width int) {
	padding := (width - utf8.RuneCountInString(text)) / 2
	//Bug- len is in bytes and
	//padding := (width - len(text)) / 2
	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}
	fmt.Println(text)
	for i := 0; i < width; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}
