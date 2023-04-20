package main

import "fmt"

func banner(text string, width int) {
	padding := (width - len(text)) / 2
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
	banner("G😄", 6)

	s := "G😄"
	fmt.Println("len:", len(s))

	g := "Go"
	fmt.Println("len:", len(g))

	//code point = rune ~=unicode
	for i, r := range s {
		fmt.Println(i, r)
	}
	// 	for i, r := range s {
	// 		fmt.Println(i, r)
	// 		if i == 0 {
	// 			fmt.Printf("%c of type %T\n", r, r)
	// 			//rune =int32

	// }
}
