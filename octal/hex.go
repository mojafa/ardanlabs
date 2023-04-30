package main

import (
	"fmt"
)

func main() {
	adams := 42
	fmt.Printf("42 as binary, %b \n", adams)
	fmt.Printf("42 as hexadecimal, %x \n", adams)

	// print these values as both binary & hexadecimal
	a, b, c, d, e, f := 0, 1, 2, 3, 4, 5
	fmt.Printf("0 as binary, %b \n", a)
	fmt.Printf("1 as binary, %b \n", b)
	fmt.Printf("2 as binary, %b \n", c)
	fmt.Printf("3 as binary, %b \n", d)
	fmt.Printf("4 as binary, %b \n", e)
	fmt.Printf("5 as binary, %b \n", f)

	fmt.Printf("0 as hexadecimal, %x \n", a)
	fmt.Printf("1 as hexadecimal, %x \n", b)
	fmt.Printf("2 as hexadecimal, %x \n", c)
	fmt.Printf("3 as hexadecimal, %x \n", d)
	fmt.Printf("4 as hexadecimal, %x \n", e)
	fmt.Printf("5 as hexadecimal, %x \n", f)

	//a better way fo doing hex
	fmt.Printf("0 as hexadecimal, %#x \n", a)
	fmt.Printf("1 as hexadecimal, %#x \n", b)
	fmt.Printf("2 as hexadecimal, %#x \n", c)
	fmt.Printf("3 as hexadecimal, %#x \n", d)
	fmt.Printf("4 as hexadecimal, %#x \n", e)
	fmt.Printf("5 as hexadecimal, %#x \n", f)

}
