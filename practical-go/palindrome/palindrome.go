package main

import "fmt"

func isPalindrome(s string) bool {
	// make it unicode aware
	rs := []rune(s)

	for i := 0; i < len(rs)/2; i++ {
		if rs[i] != rs[len(rs)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("anna"))
	fmt.Println(isPalindrome("12321"))
	fmt.Println(isPalindrome("Arewenotdrawnonwardwefewdrawnonwardtonewera"))
	fmt.Println(isPalindrome("not a palindrome"))
	fmt.Println(isPalindrome("'Are we not drawn onward, we few, drawn onward to new era? ' - Socrates Setarcos"))

}
