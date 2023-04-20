package main

import (
	"fmt"
	"unicode/utf8"
)

func banner(text string, width int) {
	// BUG len is in bytes and we need to change it to runes for the unicode emoji we added
	// padding := (width - len(text)) / 2

	//FIX
	padding := (width - utf8.RuneCountInString(text)) / 2
	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}
	fmt.Println(text)
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
}

func main() {
	banner("Go", 6)
	g := "Go"
	fmt.Println("len:", len(g))
	banner("G☺", 6)
	s := "G☺"
	fmt.Println("len:", len(s))

	//code point = rune ~=unicode
	for i, r := range s {
		fmt.Println(i, r)
	}

	//strings in GO as a duality between a seq of bytes which is utf-encoding and collection of runes/characters

	// bytes uint8, byte
	//byte int32 unicode codepoint/character
	b := s[1]
	fmt.Printf("%c of type %T\n", b, b)
	//prints out byte uint8

	for i, r := range s {
		fmt.Println(i, r)
		if i == 0 {
			fmt.Printf("%c of type %T\n", r, r)
			//prints out a rune (int32)
		}
	}
}
